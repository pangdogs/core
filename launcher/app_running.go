package launcher

import (
	"context"
	"fmt"
	"github.com/pangdogs/galaxy"
	"github.com/pangdogs/galaxy/pt"
	"github.com/pangdogs/galaxy/service"
	"github.com/pangdogs/galaxy/util"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"sync"
	"syscall"
)

func (app *App) runApp() {
	serviceConfTab := app.loadPtConfig(app.ptConfig)

	if len(app.services) <= 0 {
		for servicePtName, _ := range serviceConfTab {
			app.services = append(app.services, servicePtName)
		}
	}

	viper.SetConfigFile(app.appConfig)
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("load app config file '%s' failed, %v", app.appConfig, err))
	}

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

	for _, servicePtName := range app.services {
		servicePtConf, ok := serviceConfTab[servicePtName]
		if !ok {
			panic(fmt.Errorf("service '%s' prototype config not found", servicePtName))
		}

		wg.Add(1)
		go app.runService(ctx, &wg, servicePtName, servicePtConf)
	}

	wg.Wait()
}

func (app *App) runService(ctx context.Context, wg *sync.WaitGroup, servicePtName string, servicePtConf pt.ServiceConf) {
	defer wg.Done()

	app.newService(ctx, servicePtName, servicePtConf)

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

func (app *App) newService(ctx context.Context, servicePtName string, servicePtConf pt.ServiceConf) galaxy.Service {
	entityLib := pt.NewEntityLib()

	for entityPtName, entityPtConf := range servicePtConf.EntityTab {
		entityLib.Register(entityPtName, entityPtConf)
	}

	entityLib.Get("Daemon").New()

	serviceCtx := service.NewContext(
		service.ContextOption.Prototype(servicePtName),
		service.ContextOption.NodeID(viper.GetInt64(fmt.Sprintf("%s.NodeID", servicePtName))),
		service.ContextOption.ParentContext(ctx),
	)

	return galaxy.NewService(serviceCtx)
}

//func (app *App) newSingleton(serviceCtx core.Runtime, servicePtName string, singleton []string) {
//	singletonContext, singletonCancel := context.WithCancel(context.Background())
//
//	singletonRuntimeCtx := core.NewRuntimeContext(serviceCtx,
//		core.RuntimeContextOptionSetter.ReportError(make(chan error, 100)),
//		core.RuntimeContextOptionSetter.ParentContext(singletonContext),
//	)
//
//	singletonRuntime := core.NewRuntime(singletonRuntimeCtx,
//		core.RuntimeOptionSetter.EnableAutoRecover(viper.GetBool(fmt.Sprintf("%s.Singleton.EnableAutoRecover", servicePtName))),
//		core.RuntimeOptionSetter.ProcessQueueCapacity(viper.GetInt(fmt.Sprintf("%s.Singleton.ProcessQueueCapacity", servicePtName))),
//		core.RuntimeOptionSetter.ProcessQueueTimeout(time.Duration(viper.GetInt64(fmt.Sprintf("%s.Singleton.ProcessQueueTimeout", servicePtName)))),
//		core.RuntimeOptionSetter.GCInterval(time.Duration(viper.GetInt64(fmt.Sprintf("%s.Singleton.GCInterval", servicePtName)))),
//	)
//
//	for _, entityPtName := range servicePtConf.Singleton {
//		entityLib.Get(entityPtName).New()
//	}
//
//	singletonShutChan := singletonRuntime.Run()
//}
