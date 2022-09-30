package launcher

import (
	"gopkg.in/alecthomas/kingpin.v2"
)

// App 应用
type App struct {
	ptConfig  string
	appConfig string
	services  []string
}

// Run 运行
func (app *App) Run() {
	// 运行相关选项
	var runApp = kingpin.Command("run", "开始运行。").Default()
	var ptConfig = runApp.Flag("pt_config", "原型配置文件(*.json|*.xml)。").Default("pt.json").ExistingFile()
	var appConfig = runApp.Flag("app_config", "应用配置文件(*.json|*.xml)。").Default("app.json").ExistingFile()
	var services = runApp.Flag("services", "需要启动的服务列表。").Strings()

	// 打印相关选项
	var printInfo = kingpin.Command("print", "打印信息。").Alias("p")
	var printComp = printInfo.Command("comp", "打印所有组件。")

	switch kingpin.Parse() {
	case runApp.FullCommand():
		app.ptConfig = *ptConfig
		app.appConfig = *appConfig
		app.services = *services
		app.runApp()
		return
	case printInfo.FullCommand():
		return
	case printComp.FullCommand():
		app.printComp()
		return
	default:
		kingpin.Usage()
		return
	}
}
