package galaxy

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/pangdogs/galaxy/core"
	"github.com/pangdogs/galaxy/pt"
	"github.com/pangdogs/galaxy/util"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
	"os/signal"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"syscall"
)

type App struct {
}

func (app *App) Run() {
	var runApp = kingpin.Command("run", "开始运行。").Default()
	var ptConfFile = runApp.Flag("pt", "原型配置文件(json|xml)。").Default("pt.json").String()
	var printInfo = kingpin.Command("print", "打印信息。").Alias("p")
	var printComp = printInfo.Command("comp", "打印所有组件。")

	switch kingpin.Parse() {
	case runApp.FullCommand():
		app.runApp(*ptConfFile)
		return
	case printInfo.FullCommand():
		return
	case printComp.FullCommand():
		app.printComp()
		return
	case kingpin.HelpCommand.FullCommand():
		kingpin.Usage()
		return
	}
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

func (app *App) printComp() {
	var compPts []pt.ComponentPt

	pt.RangeCompPts(func(compPt pt.ComponentPt) bool {
		compPts = append(compPts, compPt)
		return true
	})

	sort.Slice(compPts, func(i, j int) bool {
		return strings.Compare(compPts[i].Api+compPts[i].Tag, compPts[j].Api+compPts[j].Tag) < 0
	})

	compPtsData, err := json.MarshalIndent(compPts, "", "\t")
	if err != nil {
		panic(fmt.Errorf("marshal components prototype info failed, %v", err))
	}

	fmt.Printf("%s", compPtsData)
}
