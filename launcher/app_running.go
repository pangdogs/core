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

func (app *_App) runService(ctx context.Context, serviceName string, serviceConf ServiceConf) {
	entityLib := pt.NewEntityLib()

	for entityPtName, entityPtConf := range serviceConf.EntityTab {
		entityLib.Register(entityPtName, entityPtConf.ComponentTab)
	}

	pluginBundle := plugin.NewPluginBundle()

	if app.options.ServiceInstallPlugin != nil {
		app.options.ServiceInstallPlugin(serviceName, pluginBundle)
	}

	var autoRecover bool
	var reportError chan error
	if app.options.ServiceSetupRecover != nil {
		autoRecover, reportError = app.options.ServiceSetupRecover(serviceName)
	}

	var startedCallback, stoppingCallback, stoppedCallback func(serviceCtx service.Context)
	if app.options.ServiceSetupStartedCallback != nil {
		startedCallback = app.options.ServiceSetupStartedCallback(serviceName)
	}
	if app.options.ServiceSetupStoppingCallback != nil {
		stoppingCallback = app.options.ServiceSetupStoppingCallback(serviceName)
	}
	if app.options.ServiceSetupStoppedCallback != nil {
		stoppedCallback = app.options.ServiceSetupStoppedCallback(serviceName)
	}

	service := golaxy.NewService(service.NewContext(
		service.WithContextOption{}.Context(ctx),
		service.WithContextOption{}.AutoRecover(autoRecover),
		service.WithContextOption{}.ReportError(reportError),
		service.WithContextOption{}.Name(serviceName),
		service.WithContextOption{}.EntityLib(entityLib),
		service.WithContextOption{}.PluginBundle(pluginBundle),
		service.WithContextOption{}.StartedCallback(startedCallback),
		service.WithContextOption{}.StoppingCallback(stoppingCallback),
		service.WithContextOption{}.StoppedCallback(stoppedCallback),
	))
	<-service.Run()
}
