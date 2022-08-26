package core

import "reflect"

// As 从实体提取一些需要的组件API（Component API），复合在一起直接使用，提取失败不会panic，非线程安全，例如：
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
	complexApi := Zero[T]()
	vfComplexApi := reflect.ValueOf(&complexApi).Elem()

	if vfComplexApi.Kind() != reflect.Struct {
		return Zero[T](), false
	}

	for i := 0; i < vfComplexApi.NumField(); i++ {
		vfCompApi := vfComplexApi.Field(i)

		if vfCompApi.Kind() != reflect.Interface {
			return Zero[T](), false
		}

		tfCompApi := vfCompApi.Type()
		compApiName := tfCompApi.PkgPath() + "/" + tfCompApi.Name()

		comp := entity.GetComponent(compApiName)
		if comp == nil {
			return Zero[T](), false
		}

		vfCompApi.Set(comp.getReflectValue())
	}

	return complexApi, true
}

// Cast 与As功能相同，只是提取失败时会panic，非线程安全
func Cast[T any](entity Entity) T {
	entityFace, ok := As[T](entity)
	if !ok {
		panic("cast invalid")
	}
	return entityFace
}
