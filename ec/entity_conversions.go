package ec

import (
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
//	type ComplexIface struct {
//		A
//		B
//	}
//	...
//	c, ok := As[ComplexIface](entity)
//	if ok {
//		c.MethodA()
//		c.MethodB()
//	}
//
// 注意：
//
//	1.内部逻辑有使用反射，效率较差，最好使用一次后存储转换结果重复使用。
//	2.提取完成后，实体删除或更换组件有，需要重新提取。
func As[T comparable](entity Entity) (T, bool) {
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
//	type ComplexIface struct {
//		A
//		B
//	}
//	...
//	Cast[ComplexIface](entity).MethodA()
//	Cast[ComplexIface](entity).MethodB()
//
// 注意：
//
//	1.内部逻辑有使用反射，效率较差，最好使用一次后存储转换结果重复使用。
//	2.提取完成后，实体删除或更换组件有，需要重新提取。
func Cast[T comparable](entity Entity) T {
	entityFace, ok := As[T](entity)
	if !ok {
		panic("incorrect cast")
	}
	return entityFace
}

// NewComplex 创建组件复合提取器，直接使用As()或Cast()无法检测提取后是否失效，使用提取器可以解决此问题，示例：
//
//	type A interface {
//		MethodA()
//	}
//	...
//	type B interface {
//		MethodB()
//	}
//	...
//	type ComplexIface struct {
//		A
//		B
//	}
//	...
//	cx := NewComplex[ComplexIface](nil)
//	...
//	if rv := cx.As(func(complex Complex[ComplexIface], complexIface ComplexIface) {
//		complexIface.MethodA()
//
//		if rv := complex.As(func(complex Complex[ComplexIface], complexIface ComplexIface) {
//			complexIface.MethodB()
//		}); !rv {
//			...
//		}
//	}); !rv {
//		...
//	}
//	...
//	cx.Cast(func(complex Complex[ComplexIface], complexIface ComplexIface) {
//		complexIface.MethodA()
//
//		complex.Cast(func(complex Complex[ComplexIface], complexIface ComplexIface) {
//			complexIface.MethodB()
//		})
//	})
func NewComplex[T comparable](entity Entity) Complex[T] {
	if entity == nil {
		panic("nil entity")
	}
	return &_Complex[T]{entity: entity}
}

// Complex 组件复合提取器接口
type Complex[T comparable] interface {
	// Entity 获取实体
	Entity() Entity
	// Changed 实体是否有删除或更换组件
	Changed() bool
	// As 从实体提取一些需要的组件接口，复合在一起直接使用
	As(fun func(complex Complex[T], complexIface T)) bool
	// Cast 从实体提取一些需要的组件接口，复合在一起直接使用，提取失败会panic
	Cast(fun func(complex Complex[T], complexIface T))
}

type _Complex[T comparable] struct {
	entity         Entity
	changedVersion int64
	complexIface   T
	vfComplexIface reflect.Value
}

func (c *_Complex[T]) invoke(fun func(complex Complex[T], complexIface T)) bool {
	if fun == nil {
		return false
	}

	if c.complexIface != util.Zero[T]() && !c.Changed() {
		fun(c, c.complexIface)
		return false
	}

	if !c.vfComplexIface.IsValid() {
		c.vfComplexIface = reflect.ValueOf(c.complexIface)
	}

	if !as(c.entity, c.vfComplexIface) {
		return false
	}

	c.changedVersion = c.entity.getChangedVersion()

	fun(c, c.complexIface)

	return true
}

// Entity 获取实体
func (c *_Complex[T]) Entity() Entity {
	return c.entity
}

// Changed 实体是否有删除或更换组件
func (c *_Complex[T]) Changed() bool {
	return c.changedVersion != c.entity.getChangedVersion()
}

// As 从实体提取一些需要的组件接口，复合在一起直接使用
func (c *_Complex[T]) As(fun func(complex Complex[T], complexIface T)) bool {
	return c.invoke(fun)
}

// Cast 从实体提取一些需要的组件接口，复合在一起直接使用，提取失败会panic
func (c *_Complex[T]) Cast(fun func(complex Complex[T], complexIface T)) {
	if !c.invoke(fun) {
		panic("incorrect cast")
	}
}

func as(entity Entity, vfComplexIface reflect.Value) bool {
	sb := strings.Builder{}
	sb.Grow(128)

	switch vfComplexIface.Kind() {
	case reflect.Struct:
		for i := 0; i < vfComplexIface.NumField(); i++ {
			vfCompIface := vfComplexIface.Field(i)

			if vfCompIface.Kind() != reflect.Interface {
				return false
			}

			tfCompIface := vfCompIface.Type()

			sb.Reset()
			sb.WriteString(tfCompIface.PkgPath())
			sb.WriteString("/")
			sb.WriteString(tfCompIface.Name())

			comp := entity.GetComponent(sb.String())
			if comp == nil {
				return false
			}

			vfCompIface.Set(comp.getReflectValue())
		}

		return true

	case reflect.Interface:
		tfComplexIface := vfComplexIface.Type()

		sb.Reset()
		sb.WriteString(tfComplexIface.PkgPath())
		sb.WriteString("/")
		sb.WriteString(tfComplexIface.Name())

		comp := entity.GetComponent(sb.String())
		if comp == nil {
			return false
		}

		return true

	default:
		return false
	}
}

// GetInheritor 获取实体的继承者
func GetInheritor[T any](entity Entity) T {
	return util.Cache2Iface[T](entity.getOptions().Inheritor.Cache)
}
