package launcher

// ServiceConfTab 服务原型配置表
type ServiceConfTab map[string]ServiceConf

// ServiceConf 服务原型配置
type ServiceConf struct {
	EntityTab EntityConfTab `json:"Entity"`
}

// EntityConfTab 实体原型配置表
type EntityConfTab map[string]EntityConf

// EntityConf 实体原型配置
type EntityConf struct {
	ComponentTab ComponentConfTab `json:"Component"`
	Singleton    bool             `json:"Singleton"`
}

// ComponentConfTab 组件原型配置表
type ComponentConfTab []string
