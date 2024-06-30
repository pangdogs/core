package pt

import (
	"git.golaxy.org/core/ec"
	"git.golaxy.org/core/utils/types"
)

type _CompAlias struct {
	Comp  any
	Alias string
	Fixed bool
}

// CompAlias 组件与别名，用于注册实体原型时自定义组件别名
func CompAlias(comp any, fixed bool, alias string) _CompAlias {
	return _CompAlias{
		Comp:  comp,
		Alias: alias,
		Fixed: fixed,
	}
}

// CompInterface 组件与接口，用于注册实体原型时使用接口名作为别名
func CompInterface[FACE any](comp any, fixed bool) _CompAlias {
	return _CompAlias{
		Comp:  comp,
		Alias: types.FullNameT[FACE](),
		Fixed: fixed,
	}
}

// Attribute 实体原型属性
type Attribute struct {
	Composite          any       // 实体类型
	Scope              *ec.Scope // 可访问作用域
	AwakeOnFirstAccess *bool     // 设置开启组件被首次访问时，检测并调用Awake()
}
