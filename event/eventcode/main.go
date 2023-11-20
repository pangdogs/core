// Package eventcode 使用go:generate功能，在编译前自动化生成代码
/*
	- 可以生成事件（event）与事件表（event table）辅助代码。
	- 用于生成事件辅助代码时，在事件定义代码源文件（.go）头部，添加以下注释：
		//go:generate go run kit.golaxy.org/golaxy/event/eventcode --decl_file=$GOFILE gen_event --package=$GOPACKAGE
	- 用于生成事件表辅助代码时，在事件定义代码源文件（.go）头部，添加以下注释：
		//go:generate go run kit.golaxy.org/golaxy/event/eventcode --decl_file=$GOFILE gen_eventtab --package=$GOPACKAGE --name={事件表名称}
	- 在cmd控制台中，进入事件定义代码源文件（.go）的目录，输入go generate指令即可生成代码，此外也可以使用IDE提供的go generate功能。
	- 编译本包并执行eventcode --help，可以查看命令行参数，通过参数可以调整生成的代码。
*/
package main

import (
	"fmt"
	"github.com/alecthomas/kingpin/v2"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"
)

type CommandContext struct {
	// 基础选项
	DeclFile          string
	EventRegexp       string
	PackageEventAlias string
	PackageIfaceAlias string
	FileData          []byte
	FileSet           *token.FileSet
	FileAst           *ast.File

	// 生成事件代码相关选项
	EventPackage   string
	EventDir       string
	EventDefExport bool
	EventDefAuto   bool

	// 生成事件表代码相关选项
	EventTabPackage string
	EventTabDir     string
	EventTabName    string
}

const (
	packageEventPath = "kit.golaxy.org/golaxy/event"
	packageIfacePath = "kit.golaxy.org/golaxy/util/iface"
)

func main() {
	// 基础选项
	declFile := kingpin.Flag("decl_file", "定义事件的源文件（.go）。").ExistingFile()
	eventRegexp := kingpin.Flag("event_regexp", "匹配事件定义时，使用的正则表达式。").Default("^[eE]vent.+").String()
	packageEventAlias := kingpin.Flag("package_event_alias", fmt.Sprintf("导入GOLAXY框架的`%s`包时使用的别名。", packageEventPath)).Default("event").String()
	packageIfaceAlias := kingpin.Flag("package_iface_alias", fmt.Sprintf("导入GOLAXY框架的`%s`包时使用的别名。", packageIfacePath)).Default("iface").String()

	// 生成事件代码相关选项
	eventCmd := kingpin.Command("gen_event", "通过定义的事件，生成事件辅助代码。")
	eventPackage := eventCmd.Flag("package", "生成事件辅助代码时，使用的包名。").String()
	eventDir := eventCmd.Flag("dir", "生成事件辅助代码时，输出的源文件（.go）存放的相对目录。").String()
	eventDefExport := eventCmd.Flag("default_export", "生成事件辅助代码时，发送事件的辅助代码的可见性，事件定义选项[EmitExport][EmitUnExport]可以覆盖此配置。").Default("true").String()
	eventDefAuto := eventCmd.Flag("default_auto", "生成事件辅助代码时，是否生成简化绑定事件的辅助代码，事件定义选项[EmitAuto][EmitManual]可以覆盖此配置。").Default("true").String()

	// 生成事件表代码相关选项
	eventTabCmd := kingpin.Command("gen_eventtab", "通过定义的事件，生成事件表辅助代码。")
	eventTabPackage := eventTabCmd.Flag("package", "生成事件表辅助代码，使用的包名。").String()
	eventTabDir := eventTabCmd.Flag("dir", "生成事件表辅助代码时，输出的源文件（.go）存放的相对目录。").String()
	eventTabName := eventTabCmd.Flag("name", "生成的事件表名称。").String()

	cmd := kingpin.Parse()

	ctx := &CommandContext{}
	ctx.DeclFile, _ = filepath.Abs(*declFile)
	ctx.EventRegexp = strings.TrimSpace(*eventRegexp)
	ctx.PackageEventAlias = strings.TrimSpace(*packageEventAlias)
	ctx.PackageIfaceAlias = strings.TrimSpace(*packageIfaceAlias)

	if ctx.PackageEventAlias == "" {
		panic("`gen_event --package_event_alias`设置的别名不能为空")
	}

	if ctx.PackageIfaceAlias == "" {
		panic("`gen_event --package_iface_alias`设置的别名不能为空")
	}

	switch cmd {
	case eventCmd.FullCommand():
		loadDeclFile(ctx)

		ctx.EventPackage = strings.TrimSpace(*eventPackage)
		ctx.EventDir = strings.TrimSpace(*eventDir)
		ctx.EventDefExport, _ = strconv.ParseBool(*eventDefExport)
		ctx.EventDefAuto, _ = strconv.ParseBool(*eventDefAuto)

		if ctx.EventPackage == "" {
			panic("`gen_event --package`设置的包名不能为空")
		}

		genEvent(ctx)
		return
	case eventTabCmd.FullCommand():
		loadDeclFile(ctx)

		ctx.EventTabPackage = *eventTabPackage
		ctx.EventTabDir = *eventTabDir
		ctx.EventTabName = *eventTabName

		if ctx.EventTabPackage == "" {
			panic("`gen_eventtab --package`设置的包名不能为空")
		}

		if ctx.EventTabName == "" {
			panic("`gen_eventtab --name`设置的事件表名不能为空")
		}

		genEventTab(ctx)
		return
	default:
		kingpin.Usage()
		return
	}
}

func loadDeclFile(ctx *CommandContext) {
	if filepath.Ext(ctx.DeclFile) != ".go" {
		panic("`--decl_file`设置的源文件后缀名必须为`.go`")
	}

	if ctx.EventRegexp == "" {
		panic("`--event_regexp`设置的正则表达式不能为空")
	}

	fileData, err := ioutil.ReadFile(ctx.DeclFile)
	if err != nil {
		panic(err)
	}

	fset := token.NewFileSet()

	fast, err := parser.ParseFile(fset, ctx.DeclFile, fileData, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	ctx.FileData = fileData
	ctx.FileSet = fset
	ctx.FileAst = fast
}
