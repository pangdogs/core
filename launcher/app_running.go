package launcher

import (
	"context"
	"fmt"
	"kit.golaxy.org/golaxy"
	"kit.golaxy.org/golaxy/plugin"
	"kit.golaxy.org/golaxy/pt"
	"kit.golaxy.org/golaxy/service"
	"kit.golaxy.org/golaxy/util"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"sync"
	"syscall"
)

func (app *_App) runApp(services []string, ptPath string) {
	servicePtConf := app.loadPtConfig(ptPath)

	if len(services) <= 0 {
		for service, _ := range servicePtConf {
			services = append(services, service)
		}
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

	for _, service := range services {
		serviceConf, ok := servicePtConf[service]
		if !ok {
			panic(fmt.Errorf("service '%s' pt config not found", service))
		}

		wg.Add(1)
		go func(service string, serviceConf ServiceConf) {
			defer wg.Done()
			app.runService(ctx, service, serviceConf)
		}(service, serviceConf)
	}

	wg.Wait()
}

func (app *_App) loadPtConfig(ptConfFile string) ServiceConfTab {
	switch strings.ToLower(filepath.Ext(ptConfFile)) {
	case ".json":
		loader := util.JsonLoader[ServiceConfTab]{}

		if err := loader.SetFile(ptConfFile); err != nil {
			panic(fmt.Errorf("load service pt config '%s' failed, %v", ptConfFile, err))
		}

		return loader.Get()

	case ".xml":
		loader := util.XmlLoader[ServiceConfTab]{}

		if err := loader.SetFile(ptConfFile); err != nil {
			panic(fmt.Errorf("load service pt config '%s' failed, %v", ptConfFile, err))
		}

		return loader.Get()

	default:
		panic(fmt.Errorf("load service pt config '%s' failed, file suffix invalid", ptConfFile))
	}
}

func (app *_App) runService(ctx context.Context, servicePt string, serviceConf ServiceConf) {
	entityLib := pt.NewEntityLib()

	for entityPtName, entityPtConf := range serviceConf.EntityTab {
		entityLib.Register(entityPtName, entityPtConf.ComponentTab)
	}

	pluginBundle := plugin.NewPluginBundle()

	serviceCtxOpts := []service.ContextOption{
		service.WithContextOption{}.Context(ctx),
		service.WithContextOption{}.Prototype(servicePt),
		service.WithContextOption{}.EntityLib(entityLib),
		service.WithContextOption{}.PluginBundle(pluginBundle),
	}

	if app.options.ServiceCtxInitTab != nil {
		initFunc := app.options.ServiceCtxInitTab[servicePt]
		if initFunc != nil {
			serviceCtxOpts = append(serviceCtxOpts, initFunc(entityLib, pluginBundle)...)
		}
	}

	var serviceOpts []golaxy.ServiceOption

	if app.options.ServiceInitTab != nil {
		initFunc := app.options.ServiceInitTab[servicePt]
		if initFunc != nil {
			serviceOpts = append(serviceOpts, initFunc()...)
		}
	}

	service := golaxy.NewService(service.NewContext(serviceCtxOpts...), serviceOpts...)
	<-service.Run()
}
