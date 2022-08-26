package core

// _ComponentQuery 组件查询接口
type _ComponentQuery interface {
	// GetComponent 使用名称查询组件（Component），一般情况下名称指组件接口名称，也可以自定义名称，同个名称指向多个组件时，返回首个组件，非线程安全
	GetComponent(name string) Component

	// GetComponentByID 使用组件（Component）运行时ID查询组件，非线程安全
	GetComponentByID(id uint64) Component

	// GetComponents 使用名称查询所有指向的组件（Component），非线程安全
	GetComponents(name string) []Component

	// RangeComponents 遍历所有组件，非线程安全
	RangeComponents(fun func(component Component) bool)
}

// _ComponentMgr 组件管理器接口
type _ComponentMgr interface {
	_ComponentQuery

	// AddComponents 使用同个名称添加多个组件（Component），一般情况下名称指组件接口名称，也可以自定义名称，非线程安全
	AddComponents(name string, components []Component) error

	// AddComponent 添加单个组件（Component），因为同个名称可以指向多个组件，所有名称指向的组件已存在时，不会返回错误，非线程安全
	AddComponent(name string, component Component) error

	// RemoveComponent 删除名称指向的组件（Component），会删除同个名称指向的多个组件，非线程安全
	RemoveComponent(name string)

	// RemoveComponentByID 使用组件（Component）运行时ID删除组件，非线程安全
	RemoveComponentByID(id uint64)
}
