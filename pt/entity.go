package pt

import (
	"fmt"
	"git.golaxy.org/core/ec"
	"git.golaxy.org/core/internal/exception"
	"git.golaxy.org/core/util/iface"
	"git.golaxy.org/core/util/option"
	"reflect"
)

// CompInfo 组件信息
type CompInfo struct {
	PT    ComponentPT // 原型
	Alias string      // 别名
}

// EntityPT 实体原型
type EntityPT struct {
	Prototype          string       // 实体原型名称
	CompositeRT        reflect.Type // 扩展者反射类型
	Scope              *ec.Scope    // 可访问作用域
	AwakeOnFirstAccess *bool        // 设置开启组件被首次访问时，检测并调用Awake()
	Components         []CompInfo   // 组件信息
}

// Construct 创建实体
func (pt EntityPT) Construct(settings ...option.Setting[ec.EntityOptions]) ec.Entity {
	options := option.Make(ec.With.Default())
	if pt.CompositeRT != nil {
		options.CompositeFace = iface.MakeFace(reflect.New(pt.CompositeRT).Interface().(ec.Entity))
	}
	if pt.Scope != nil {
		options.Scope = *pt.Scope
	}
	if pt.AwakeOnFirstAccess != nil {
		options.AwakeOnFirstAccess = *pt.AwakeOnFirstAccess
	}
	options = option.Append(options, settings...)
	options.Prototype = pt.Prototype

	return pt.assemble(ec.UnsafeNewEntity(options))
}

func (pt EntityPT) assemble(entity ec.Entity) ec.Entity {
	if entity == nil {
		panic(fmt.Errorf("%w: %w: entity is nil", ErrPt, exception.ErrArgs))
	}

	for i := range pt.Components {
		compInfo := &pt.Components[i]

		comp := compInfo.PT.Construct()

		if err := entity.AddComponent(compInfo.Alias, comp); err != nil {
			panic(fmt.Errorf("%w: %w", ErrPt, err))
		}
	}

	return entity
}
