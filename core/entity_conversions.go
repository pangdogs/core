package core

import (
	"reflect"
	"strings"
)

// As 从实体提取一些需要的组件接口（Component Interface），复合在一起直接使用，提取失败不会panic，非线程安全，例如：
//	type A interface {
//		TestA()
//	}
//	...
//	type B interface {
//		TestB()
//	}
//	...
//	type Complex struct {
//		api.A
//		api.B
//	}
//	...
//	As[Complex](entity).TestA()
//	As[Complex](entity).TestB()
// 注意提取后从实体删除或更换组件后，需要重新提取
func As[T any](entity Entity) (T, bool) {
	complexIface := Zero[T]()
	vfComplexIface := reflect.ValueOf(&complexIface).Elem()

	sb := strings.Builder{}
	sb.Grow(128)

	switch vfComplexIface.Kind() {
	case reflect.Struct:
		for i := 0; i < vfComplexIface.NumField(); i++ {
			vfCompIface := vfComplexIface.Field(i)

			if vfCompIface.Kind() != reflect.Interface {
				return Zero[T](), false
			}

			tfCompIface := vfCompIface.Type()

			sb.Reset()
			sb.WriteString(tfCompIface.PkgPath())
			sb.WriteString("/")
			sb.WriteString(tfCompIface.Name())

			comp := entity.GetComponent(sb.String())
			if comp == nil {
				return Zero[T](), false
			}

			vfCompIface.Set(comp.getReflectValue())
		}

		return complexIface, true

	case reflect.Interface:
		tfComplexIface := vfComplexIface.Type()

		sb.Reset()
		sb.WriteString(tfComplexIface.PkgPath())
		sb.WriteString("/")
		sb.WriteString(tfComplexIface.Name())

		comp := entity.GetComponent(sb.String())
		if comp == nil {
			return Zero[T](), false
		}

		return complexIface, true

	default:
		return Zero[T](), false
	}
}

// Cast 与As功能相同，只是提取失败时会panic，非线程安全
func Cast[T any](entity Entity) T {
	entityFace, ok := As[T](entity)
	if !ok {
		panic("cast invalid")
	}
	return entityFace
}
