package launcher

import (
	"context"
	"fmt"
	"github.com/golaxy-kit/golaxy/pt"
	"github.com/golaxy-kit/golaxy/service"
	"github.com/golaxy-kit/golaxy/util"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"sync"
	"syscall"
	"time"
)

func (app *App) runApp() {
	serviceConfTab := app.loadPtConfig(app.ptConfig)

	if len(app.services) <= 0 {
		for service, _ := range serviceConfTab {
			app.services = append(app.services, service)
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

	for _, service := range app.services {
		serviceConf, ok := serviceConfTab[service]
		if !ok {
			panic(fmt.Errorf("service '%s' prototype config not found", service))
		}

		wg.Add(1)
		go app.runService(ctx, &wg, service, serviceConf)
	}

	wg.Wait()
}

func (app *App) runService(ctx context.Context, wg *sync.WaitGroup, serviceName string, serviceConf ServiceConf) {
	defer wg.Done()

	app.newService(ctx, serviceName, serviceConf)

}

func (app *App) loadPtConfig(ptConfFile string) ServiceConfTab {
	switch strings.ToLower(filepath.Ext(ptConfFile)) {
	case ".json":
		loader := util.JsonLoader[ServiceConfTab]{}

		if err := loader.SetFile(ptConfFile); err != nil {
			panic(fmt.Errorf("load prototype config file '%s' failed, %v", ptConfFile, err))
		}

		return loader.Get()

	case ".xml":
		loader := util.XmlLoader[ServiceConfTab]{}

		if err := loader.SetFile(ptConfFile); err != nil {
			panic(fmt.Errorf("load prototype config file '%s' failed, %v", ptConfFile, err))
		}

		return loader.Get()

	default:
		panic(fmt.Errorf("load prototype config file '%s' failed, file suffix invalid", ptConfFile))
	}
}

func (app *App) newService(ctx context.Context, servicePtName string, servicePtConf ServiceConf) galaxy.Service {
	entityLib := pt.NewEntityLib()

	for entityPtName, entityPtConf := range servicePtConf.EntityTab {
		entityLib.Register(entityPtName, entityPtConf)
	}

	entityLib.Get("Daemon").New()

	serviceCtx := service.NewContext(
		service.WithContextOption.Prototype(servicePtName),
		service.WithContextOption.NodeID(viper.GetInt64(fmt.Sprintf("%s.NodeID", servicePtName))),
		service.WithContextOption.Context(ctx),
	)

	return galaxy.NewService(serviceCtx)
	return nil
}

func (app *App) newSingleton(serviceCtx core.Runtime, servicePtName string, singleton []string) {
	singletonContext, singletonCancel := context.WithCancel(context.Background())

	singletonRuntimeCtx := core.NewRuntimeContext(serviceCtx,
		core.RuntimeContextOptionSetter.ReportError(make(chan error, 100)),
		core.RuntimeContextOptionSetter.Context(singletonContext),
	)

	singletonRuntime := core.NewRuntime(singletonRuntimeCtx,
		core.WithRuntimeOption.AutoRecover(viper.GetBool(fmt.Sprintf("%s.Singleton.AutoRecover", servicePtName))),
		core.WithRuntimeOption.ProcessQueueCapacity(viper.GetInt(fmt.Sprintf("%s.Singleton.ProcessQueueCapacity", servicePtName))),
		core.WithRuntimeOption.ProcessQueueTimeout(time.Duration(viper.GetInt64(fmt.Sprintf("%s.Singleton.ProcessQueueTimeout", servicePtName)))),
		core.WithRuntimeOption.GCInterval(time.Duration(viper.GetInt64(fmt.Sprintf("%s.Singleton.GCInterval", servicePtName)))),
	)

	for _, entityPtName := range servicePtConf.Singleton {
		entityLib.Get(entityPtName).New()
	}

	singletonShutChan := singletonRuntime.Run()
}
