package entitylib

import (
	"fmt"
	"github.com/pangdogs/galaxy/complib"
	"github.com/pangdogs/galaxy/core"
)

type IEntityLib interface {
	Register(prototype string, compTags []string)
	New(prototype string) core.Entity
	Range(fun func(info EntityInfo) bool)
}

type EntityInfo struct {
	Prototype string
	CompTags  []string
}

type EntityLib struct {
	entityInfoMap map[string]EntityInfo
}

func (lib *EntityLib) init() {
	if lib.entityInfoMap == nil {
		lib.entityInfoMap = map[string]EntityInfo{}
	}
}

func (lib *EntityLib) Register(prototype string, compTags []string) {
	lib.init()

	if prototype == "" {
		panic("nil prototype")
	}

	_, ok := lib.entityInfoMap[prototype]
	if ok {
		panic(fmt.Errorf("repeated register entity '%s' invalid", prototype))
	}

	lib.entityInfoMap[prototype] = EntityInfo{
		Prototype: prototype,
		CompTags:  compTags,
	}
}

func (lib *EntityLib) New(prototype string) core.Entity {
	lib.init()

	info, ok := lib.entityInfoMap[prototype]
	if !ok {
		panic(fmt.Errorf("entity '%s' not registered invalid", prototype))
	}

	entity := core.NewEntity()

	for i := range info.CompTags {
		api, comp := complib.New(info.CompTags[i])
		entity.AddComponent(api, comp)
	}

	return entity
}

func (lib *EntityLib) Range(fun func(info EntityInfo) bool) {
	lib.init()

	if fun == nil {
		return
	}

	for _, info := range lib.entityInfoMap {
		if !fun(info) {
			return
		}
	}
}
