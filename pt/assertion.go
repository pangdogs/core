package pt

import (
	"fmt"
	"git.golaxy.org/core/ec"
	"git.golaxy.org/core/internal/exception"
	"git.golaxy.org/core/util/types"
	"reflect"
	"strings"
)

// As 从实体提取一些需要的组件接口，复合在一起直接使用
/*
示例：
	type A interface {
		MethodA()
	}
	...
	type B interface {
		MethodB()
	}
	...
	type CompositeAB struct {
		A
		B
	}
	...
	v, ok := As[CompositeAB](entity)
	if ok {
		v.MethodA()
		v.MethodB()
	}

注意：
	1.内部逻辑有使用反射，为了提高性能，可以使用一次后存储转换结果重复使用。
	2.实体更新组件后，需要重新提取。
*/
func As[T comparable](entity ec.Entity) (T, bool) {
	iface := types.Zero[T]()
	vfIface := reflect.ValueOf(&iface).Elem()

	if !as(entity, vfIface) {
		return types.Zero[T](), false
	}

	return iface, true
}

// Cast 从实体提取一些需要的组件接口，复合在一起直接使用，提取失败会panic
/*
示例：
	type A interface {
		MethodA()
	}
	...
	type B interface {
		MethodB()
	}
	...
	type CompositeAB struct {
		A
		B
	}
	...
	Cast[CompositeAB](entity).MethodA()
	Cast[CompositeAB](entity).MethodB()

注意：
	1.内部逻辑有使用反射，为了提高性能，可以使用一次后存储转换结果重复使用。
	2.实体更新组件后，需要重新提取。
*/
func Cast[T comparable](entity ec.Entity) T {
	iface, ok := As[T](entity)
	if !ok {
		panic(fmt.Errorf("%w: incorrect cast", ErrPt))
	}
	return iface
}

// Compose 创建组件复合器
func Compose[T comparable](entity ec.Entity) *Composite[T] {
	if entity == nil {
		panic(fmt.Errorf("%w: %w: entity is nil", ErrPt, exception.ErrArgs))
	}
	return &Composite[T]{
		entity: entity,
	}
}

// Composite 组件复合器，直接使用As()或Cast()时，无法检测提取后实体是否又更新组件，使用复合器可以解决此问题
/*
示例：
	type A interface {
		MethodA()
	}
	...
	type B interface {
		MethodB()
	}
	...
	type CompositeAB struct {
		A
		B
	}
	...
	cx := Compose[CompositeAB](entity)
	...
	if v, ok := cx.As(); ok {
		v.MethodA()
		v.MethodB()
		...
		entity.AddComponent(comp)
		...
		if v, ok := cx.As(); ok {
			v.MethodA()
			v.MethodB()
		}
	}
	...
	cx.Cast().MethodA()
	cx.Cast().MethodB()
	...
	entity.AddComponent(comp)
	...
	cx.Cast().MethodA()
	cx.Cast().MethodB()
*/
type Composite[T comparable] struct {
	entity  ec.Entity
	version int32
	iface   T
}

// Entity 实体
func (c *Composite[T]) Entity() ec.Entity {
	if c.entity == nil {
		panic(fmt.Errorf("%w: setting entity is nil", ErrPt))
	}
	return c.entity
}

// Changed 实体是否已更新组件
func (c *Composite[T]) Changed() bool {
	if c.entity == nil {
		panic(fmt.Errorf("%w: setting entity is nil", ErrPt))
	}
	return c.version != ec.UnsafeEntity(c.entity).GetVersion()
}

// As 从实体提取一些需要的组件接口，复合在一起直接使用（实体更新组件后，会自动重新提取）
func (c *Composite[T]) As() (T, bool) {
	if c.entity == nil {
		panic(fmt.Errorf("%w: setting entity is nil", ErrPt))
	}

	if c.iface != types.Zero[T]() && !c.Changed() {
		return c.iface, true
	}

	if !as(c.entity, reflect.ValueOf(c.iface)) {
		return types.Zero[T](), false
	}

	c.version = ec.UnsafeEntity(c.entity).GetVersion()

	return c.iface, true
}

// Cast 从实体提取一些需要的组件接口，复合在一起直接使用，提取失败会panic（实体更新组件后，会自动重新提取）
func (c *Composite[T]) Cast() T {
	iface, ok := c.As()
	if !ok {
		panic(fmt.Errorf("%w: incorrect cast", ErrPt))
	}
	return iface
}

func as(entity ec.Entity, vfIface reflect.Value) bool {
	if entity == nil {
		return false
	}

	sb := strings.Builder{}
	sb.Grow(128)

	switch vfIface.Kind() {
	case reflect.Struct:
		for i := 0; i < vfIface.NumField(); i++ {
			vfField := vfIface.Field(i)
			tfField := vfField.Type()

			switch vfField.Kind() {
			case reflect.Pointer:
				tfField = tfField.Elem()
				break
			case reflect.Interface:
				break
			default:
				return false
			}

			sb.Reset()
			types.WriteTypeFullName(&sb, tfField)

			comp := entity.GetComponent(sb.String())
			if comp == nil {
				return false
			}

			vfField.Set(ec.UnsafeComponent(comp).GetReflected())
		}

		return true

	case reflect.Interface:
		tfIface := vfIface.Type()

		sb.Reset()
		types.WriteTypeFullName(&sb, tfIface)

		comp := entity.GetComponent(sb.String())
		if comp == nil {
			return false
		}

		vfIface.Set(ec.UnsafeComponent(comp).GetReflected())

		return true

	default:
		return false
	}
}
