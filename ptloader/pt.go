package ptloader

// ServicePts 所有的服务原型（Service Prototype）配置表
type ServicePts map[string]ServicePt

// ServicePt 服务原型配置
type ServicePt struct {
	Entity    EntityPts
	Singleton []string
}

// EntityPts 所有的实体原型（Entity Prototype）配置表
type EntityPts map[string]EntityPt

// EntityPt 实体原型配置
type EntityPt ComponentPts

// ComponentPts 所有的组件（Component Prototype）原型配置表
type ComponentPts []string
