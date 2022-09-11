package galaxy

import (
	"gopkg.in/alecthomas/kingpin.v2"
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
	default:
		kingpin.Usage()
		return
	}
}
