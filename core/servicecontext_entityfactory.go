package core

type EntityFactory interface {
	NewEntity(params EntityParams) Entity
	GetSingleton(prototype string) Entity
}

func (servCtx *_ServiceContextBehavior) NewEntity(params EntityParams) Entity {

}

func (servCtx *_ServiceContextBehavior) GetSingleton(prototype string) Entity {

}
