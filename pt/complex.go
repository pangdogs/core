package pt

import (
	"kit.golaxy.org/golaxy/ec"
	"kit.golaxy.org/golaxy/util"
	"reflect"
	"strings"
)

// As 从实体提取一些需要的组件接口，复合在一起直接使用，示例：
//
//	type A interface {
//		MethodA()
//	}
//	...
//	type B interface {
//		MethodB()
//	}
//	...
//	type ComplexAB struct {
//		A
//		B
//	}
//	...
//	v, ok := As[ComplexAB](entity)
//	if ok {
//		v.MethodA()
//		v.MethodB()
//	}
//
// 注意：
//
//	1.内部逻辑有使用反射，为了提高性能，可以使用一次后存储转换结果重复使用。
//	2.实体更新组件后，需要重新提取。
func As[T comparable](entity ec.Entity) (T, bool) {
	complexIface := util.Zero[T]()
	vfComplexIface := reflect.ValueOf(&complexIface).Elem()

	if !as(entity, vfComplexIface) {
		return util.Zero[T](), false
	}

	return complexIface, true
}

// Cast 从实体提取一些需要的组件接口，复合在一起直接使用，提取失败会panic，示例：
//
//	type A interface {
//		MethodA()
//	}
//	...
//	type B interface {
//		MethodB()
//	}
//	...
//	type ComplexAB struct {
//		A
//		B
//	}
//	...
//	Cast[ComplexAB](entity).MethodA()
//	Cast[ComplexAB](entity).MethodB()
//
// 注意：
//
//	1.内部逻辑有使用反射，为了提高性能，可以使用一次后存储转换结果重复使用。
//	2.实体更新组件后，需要重新提取。
func Cast[T comparable](entity ec.Entity) T {
	entityFace, ok := As[T](entity)
	if !ok {
		panic("incorrect cast")
	}
	return entityFace
}

// Complex 创建组件复合提取器，直接使用As()或Cast()时，无法检测提取后实体是否又更新组件，使用提取器可以解决此问题，示例：
//
//	type A interface {
//		MethodA()
//	}
//	...
//	type B interface {
//		MethodB()
//	}
//	...
//	type ComplexAB struct {
//		A
//		B
//	}
//	...
//	cx := Complex[ComplexAB]{Entity: entity}
//	...
//	if v, ok := cx.As(); ok {
//		v.MethodA()
//		v.MethodB()
//		...
//		entity.AddComponent(comp)
//		...
//		if v, ok := cx.As(); ok {
//			v.MethodA()
//			v.MethodB()
//		}
//	}
//	...
//	cx.Cast().MethodA()
//	cx.Cast().MethodB()
//	...
//	entity.AddComponent(comp)
//	...
//	cx.Cast().MethodA()
//	cx.Cast().MethodB()
type Complex[T comparable] struct {
	Entity         ec.Entity
	changedVersion int64
	iface          T
}

// Clone 克隆
func (c Complex[T]) Clone() *Complex[T] {
	return &c
}

// Changed 实体是否已更新组件
func (c *Complex[T]) Changed() bool {
	if c.Entity == nil {
		return false
	}
	return c.changedVersion != ec.UnsafeEntity(c.Entity).GetChangedVersion()
}

// As 从实体提取一些需要的组件接口，复合在一起直接使用（实体更新组件后，会自动重新提取）
func (c *Complex[T]) As() (T, bool) {
	if c.Entity == nil {
		return util.Zero[T](), false
	}

	if c.iface != util.Zero[T]() && !c.Changed() {
		return c.iface, true
	}

	if !as(c.Entity, reflect.ValueOf(c.iface)) {
		return util.Zero[T](), false
	}

	c.changedVersion = ec.UnsafeEntity(c.Entity).GetChangedVersion()

	return c.iface, true
}

// Cast 从实体提取一些需要的组件接口，复合在一起直接使用，提取失败会panic（实体更新组件后，会自动重新提取）
func (c *Complex[T]) Cast() T {
	iface, ok := c.As()
	if !ok {
		panic("incorrect cast")
	}
	return iface
}

func as(entity ec.Entity, vfIface reflect.Value) bool {
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
			util.WriteFullName(&sb, tfCompIface)

			comp := entity.GetComponent(sb.String())
			if comp == nil {
				return false
			}

			vfCompIface.Set(ec.UnsafeComponent(comp).GetReflectValue())
		}

		return true

	case reflect.Interface:
		tfComplexIface := vfIface.Type()

		sb.Reset()
		util.WriteFullName(&sb, tfComplexIface)

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
