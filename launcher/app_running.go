package launcher

import (
	"context"
	"fmt"
	"kit.golaxy.org/golaxy"
	"kit.golaxy.org/golaxy/plugin"
	"kit.golaxy.org/golaxy/pt"
	"kit.golaxy.org/golaxy/service"
	"kit.golaxy.org/golaxy/util/config"
	"kit.golaxy.org/golaxy/util/option"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"sync"
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

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, app.options.QuitSignals...)

	go func() {
		<-sigChan
		cancel()
	}()

	for _, service := range services {
		serviceConf, ok := servicePtConf[service]
		if !ok {
			panic(fmt.Errorf("service %q pt config not found", service))
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
		loader := config.JsonLoader[ServiceConfTab]{}

		if err := loader.SetFile(ptConfFile); err != nil {
			panic(fmt.Errorf("load service pt config %q failed, %v", ptConfFile, err))
		}

		return loader.Get()

	case ".xml":
		loader := config.XmlLoader[ServiceConfTab]{}

		if err := loader.SetFile(ptConfFile); err != nil {
			panic(fmt.Errorf("load service pt config %q failed, %v", ptConfFile, err))
		}

		return loader.Get()

	default:
		panic(fmt.Errorf("load service pt config %q failed, file suffix invalid", ptConfFile))
	}
}

func (app *_App) runService(ctx context.Context, serviceName string, serviceConf ServiceConf) {
	entityLib := pt.NewEntityLib()

	for entityPtName, entityPtConf := range serviceConf.EntityTab {
		entityLib.Register(entityPtName, entityPtConf.ComponentTab)
	}

	pluginBundle := plugin.NewPluginBundle()

	serviceCtxOpts := []option.Setting[service.ContextOptions]{
		service.Option{}.Context(ctx),
		service.Option{}.Name(serviceName),
		service.Option{}.EntityLib(entityLib),
		service.Option{}.PluginBundle(pluginBundle),
	}

	for _, ctor := range app.options.ServiceCtxCtors {
		serviceCtxOpts = append(serviceCtxOpts, ctor.Exec(serviceName, entityLib, pluginBundle)...)
	}

	var serviceOpts []option.Setting[golaxy.ServiceOptions]

	for _, ctor := range app.options.ServiceCtors {
		serviceOpts = append(serviceOpts, ctor.Exec(serviceName)...)
	}

	<-golaxy.NewService(service.NewContext(serviceCtxOpts...), serviceOpts...).Run()
}
