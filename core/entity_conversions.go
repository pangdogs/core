package core

import "reflect"

func As[T any](entity Entity) (T, bool) {
	entityFace := Zero[T]()
	vfEntityFace := reflect.ValueOf(&entityFace).Elem()

	if vfEntityFace.Kind() != reflect.Struct {
		return Zero[T](), false
	}

	for i := 0; i < vfEntityFace.NumField(); i++ {
		vfCompFace := vfEntityFace.Field(i)

		if vfCompFace.Kind() != reflect.Interface {
			return Zero[T](), false
		}

		tfCompFace := vfCompFace.Type()
		ok := false

		entity.RangeComponents(func(comp Component) bool {
			vfComp := comp.getReflectValue()

			if vfComp.Type().Implements(tfCompFace) {
				vfCompFace.Set(vfComp)
				ok = true
				return false
			}

			return true
		})

		if !ok {
			return Zero[T](), false
		}
	}

	return entityFace, true
}

func Cast[T any](entity Entity) T {
	entityFace := Zero[T]()
	vfEntityFace := reflect.ValueOf(&entityFace).Elem()

	if vfEntityFace.Kind() != reflect.Struct {
		panic("ret not struct")
	}

	for i := 0; i < vfEntityFace.NumField(); i++ {
		vfCompFace := vfEntityFace.Field(i)

		if vfCompFace.Kind() != reflect.Interface {
			panic("ret field not interface")
		}

		tfCompFace := vfCompFace.Type()
		ok := false

		entity.RangeComponents(func(comp Component) bool {
			vfComp := comp.getReflectValue()

			if vfComp.Type().Implements(tfCompFace) {
				vfCompFace.Set(vfComp)
				ok = true
				return false
			}

			return true
		})

		if !ok {
			panic("ret field not matching")
		}
	}

	return entityFace
}
