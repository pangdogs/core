package pt

// ServiceConfTab 所有的服务原型（Service Prototype）配置表
type ServiceConfTab map[string]ServiceConf

// ServiceConf 服务原型配置
type ServiceConf struct {
	Entity    EntityConfTab
	Singleton []string
}

// EntityConfTab 所有的实体原型（Entity Prototype）配置表
type EntityConfTab map[string]EntityConf

// EntityConf 实体原型配置
type EntityConf ComponentConfTab

// ComponentConfTab 所有的组件（Component Prototype）原型配置表
type ComponentConfTab []string
