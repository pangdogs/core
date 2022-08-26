package core

type _EntityFactory interface {
	NewEntity(params EntityOptions) Entity
	GetSingleton(prototype string) Entity
}

// NewEntity ...
func (servCtx *_ServiceContextBehavior) NewEntity(params EntityParams) Entity {
	return nil
}

// GetSingleton ...
func (servCtx *_ServiceContextBehavior) GetSingleton(prototype string) Entity {
	return nil
}
