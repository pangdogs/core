package ec

import (
	"github.com/galaxy-kit/galaxy-go/util"
	"reflect"
	"strings"
)

// As 从实体提取一些需要的组件接口，复合在一起直接使用，提取失败不会panic，例如：
//
//	type A interface {
//		TestA()
//	}
//	...
//	type B interface {
//		TestB()
//	}
//	...
//	type Complex struct {
//		A
//		B
//	}
//	...
//	c, ok := As[Complex](entity)
//	if ok {
//		c.TestA()
//		c.TestB()
//	}
//
// 注意提取完成后，实体又删除或更换组件，需要重新提取。
func As[T any](entity Entity) (T, bool) {
	complexIface := util.Zero[T]()
	vfComplexIface := reflect.ValueOf(&complexIface).Elem()

	sb := strings.Builder{}
	sb.Grow(128)

	switch vfComplexIface.Kind() {
	case reflect.Struct:
		for i := 0; i < vfComplexIface.NumField(); i++ {
			vfCompIface := vfComplexIface.Field(i)

			if vfCompIface.Kind() != reflect.Interface {
				return util.Zero[T](), false
			}

			tfCompIface := vfCompIface.Type()

			sb.Reset()
			sb.WriteString(tfCompIface.PkgPath())
			sb.WriteString("/")
			sb.WriteString(tfCompIface.Name())

			comp := entity.GetComponent(sb.String())
			if comp == nil {
				return util.Zero[T](), false
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
			return util.Zero[T](), false
		}

		return complexIface, true

	default:
		return util.Zero[T](), false
	}
}

// Cast 从实体提取一些需要的组件接口，复合在一起直接使用，提取失败会panic，例如：
//
//	type A interface {
//		TestA()
//	}
//	...
//	type B interface {
//		TestB()
//	}
//	...
//	type Complex struct {
//		A
//		B
//	}
//	...
//	Cast[Complex](entity).TestA()
//	Cast[Complex](entity).TestB()
//
// 注意提取完成后，实体又删除或更换组件，需要重新提取。
func Cast[T any](entity Entity) T {
	entityFace, ok := As[T](entity)
	if !ok {
		panic("incorrect cast")
	}
	return entityFace
}

// GetInheritor 获取实体的继承者
func GetInheritor[T any](entity Entity) T {
	return util.Cache2Iface[T](entity.getOptions().Inheritor.Cache)
}
