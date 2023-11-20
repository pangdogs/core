package pt

import (
	"fmt"
	"kit.golaxy.org/golaxy/ec"
	"kit.golaxy.org/golaxy/util/types"
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
	compositeIface := types.Zero[T]()
	vfCompositeIface := reflect.ValueOf(&compositeIface).Elem()

	if !as(entity, vfCompositeIface) {
		return types.Zero[T](), false
	}

	return compositeIface, true
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
	entityFace, ok := As[T](entity)
	if !ok {
		panic(fmt.Errorf("%w: incorrect cast", ErrPt))
	}
	return entityFace
}

// Composite 创建组件复合提取器，直接使用As()或Cast()时，无法检测提取后实体是否又更新组件，使用提取器可以解决此问题
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
	cx := Composite[CompositeAB]{Entity: entity}
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
	Entity  ec.Entity
	version int32
	iface   T
}

// Clone 克隆
func (c Composite[T]) Clone() *Composite[T] {
	return &c
}

// Changed 实体是否已更新组件
func (c *Composite[T]) Changed() bool {
	if c.Entity == nil {
		return false
	}
	return c.version != ec.UnsafeEntity(c.Entity).GetVersion()
}

// As 从实体提取一些需要的组件接口，复合在一起直接使用（实体更新组件后，会自动重新提取）
func (c *Composite[T]) As() (T, bool) {
	if c.Entity == nil {
		return types.Zero[T](), false
	}

	if c.iface != types.Zero[T]() && !c.Changed() {
		return c.iface, true
	}

	if !as(c.Entity, reflect.ValueOf(c.iface)) {
		return types.Zero[T](), false
	}

	c.version = ec.UnsafeEntity(c.Entity).GetVersion()

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
			vfCompIface := vfIface.Field(i)

			if vfCompIface.Kind() != reflect.Interface {
				return false
			}

			tfCompIface := vfCompIface.Type()

			sb.Reset()
			types.WriteTypeFullName(&sb, tfCompIface)

			comp := entity.GetComponent(sb.String())
			if comp == nil {
				return false
			}

			vfCompIface.Set(ec.UnsafeComponent(comp).GetReflectValue())
		}

		return true

	case reflect.Interface:
		tfCompositeIface := vfIface.Type()

		sb.Reset()
		types.WriteTypeFullName(&sb, tfCompositeIface)

		comp := entity.GetComponent(sb.String())
		if comp == nil {
			return false
		}

		vfIface.Set(ec.UnsafeComponent(comp).GetReflectValue())

		return true

	default:
		return false
	}
}
