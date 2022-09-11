package galaxy

import (
	"context"
	"fmt"
	"github.com/pangdogs/galaxy/core"
	"github.com/pangdogs/galaxy/pt"
	"github.com/pangdogs/galaxy/util"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"sync"
	"syscall"
)

func (app *App) runApp(ptConfFile string) {
	serviceConfTab := app.loadPtConfig(ptConfFile)

	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, os.Kill)

	go func() {
		select {
		case <-c:
			cancel()
		}
	}()

	for servicePtName, servicePtConf := range serviceConfTab {
		wg.Add(1)
		go app.runService(ctx, &wg, servicePtName, servicePtConf)
	}

	wg.Wait()
}

func (app *App) runService(ctx context.Context, wg *sync.WaitGroup, servicePtName string, servicePtConf pt.ServiceConf) {
	defer wg.Done()

	entityLib := pt.NewEntityLib()

	for entityPtName, entityPtConf := range servicePtConf.Entity {
		entityLib.Register(entityPtName, entityPtConf)
	}

	core.NewServiceContext(ctx,
		core.ServiceContextOptionSetter.Prototype(servicePtName),
		core.ServiceContextOptionSetter.NodeID(10),
	)

}

func (app *App) loadPtConfig(ptConfFile string) pt.ServiceConfTab {
	switch strings.ToLower(filepath.Ext(ptConfFile)) {
	case ".json":
		loader := util.JsonLoader[pt.ServiceConfTab]{}

		if err := loader.SetFile(ptConfFile); err != nil {
			panic(fmt.Errorf("load prototype config file '%s' failed, %v", ptConfFile, err))
		}

		return loader.Get()

	case ".xml":
		loader := util.XmlLoader[pt.ServiceConfTab]{}

		if err := loader.SetFile(ptConfFile); err != nil {
			panic(fmt.Errorf("load prototype config file '%s' failed, %v", ptConfFile, err))
		}

		return loader.Get()

	default:
		panic(fmt.Errorf("load prototype config file '%s' failed, file suffix invalid", ptConfFile))
	}
}
