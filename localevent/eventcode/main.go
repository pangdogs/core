// Package eventcode 本地事件辅助代码生成器，使用Golang的代码生成功能（go generate）生成本地事件相关辅助代码。
//
//   - 可以生成发送事件（emit event）与事件表（event table）辅助代码。
//
//   - 用于生成发送事件辅助代码时，在事件定义代码源文件（*.go）头部，添加以下注释：
//     `go:generate go run kit.golaxy.org/golaxy/localevent/eventcode --decl_file=$GOFILE gen_emit --package=$GOPACKAGE`
//
//   - 用于生成事件表辅助代码时，在事件定义代码源文件（*.go）头部，添加以下注释：
//     `go:generate go run kit.golaxy.org/golaxy/localevent/eventcode --decl_file=$GOFILE gen_eventtab --package=$GOPACKAGE --name=XXXEventTab`
//
//   - 需要生成事件辅助代码时，在Cmd控制台中，定位到事件定义代码源文件（*.go）的路径下，输入`go generate`指令即可，也可以使用IDE提供的go generate功能。
//
//   - 本包可以编译并执行`eventcode --help`查看命令行参数，按需求调整参数改变生成的代码。
package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"gopkg.in/alecthomas/kingpin.v2"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"
)

type _CommandContext struct {
	// 基础选项
	DeclFile          string
	EventRegexp       string
	EventPackageAlias string
	FileData          []byte
	FileSet           *token.FileSet
	FileAst           *ast.File

	// 生成发送事件代码相关选项
	EmitPackage   string
	EmitDir       string
	EmitDefExport bool

	// 生成事件表代码相关选项
	EventTabPackage           string
	EventTabDir               string
	EventTabName              string
	EventTabDefEventRecursion string
}

func main() {
	// 基础选项
	declFile := kingpin.Flag("decl_file", "定义事件的源文件（*.go）。").ExistingFile()
	eventRegexp := kingpin.Flag("event_regexp", "匹配事件定义时使用的正则表达式。").Default("^[eE]vent.+").String()
	eventPackageAlias := kingpin.Flag("event_package_alias", "导入GOLAXY框架的`kit.golaxy.org/golaxy/localevent`包时使用的别名。").Default("localevent").String()

	// 生成发送事件代码相关选项
	emitCmd := kingpin.Command("gen_emit", "通过定义的事件生成发送事件代码。")
	emitPackage := emitCmd.Flag("package", "生成的发送事件代码时使用的包名。").String()
	emitDir := emitCmd.Flag("dir", "生成的发送事件代码产生的源文件（*.go）存放的相对目录。").String()
	emitDefExport := emitCmd.Flag("default_export", "生成的发送事件代码默认的可见性，事件选项[EmitExport][EmitUnExport]可以覆盖此配置。").Default("true").String()

	// 生成事件表代码相关选项
	eventTabCmd := kingpin.Command("gen_eventtab", "通过定义的事件生成事件表代码。")
	eventTabPackage := eventTabCmd.Flag("package", "生成的事件表代码使用的包名。").String()
	eventTabDir := eventTabCmd.Flag("dir", "生成的事件表产生的源文件（*.go）存放的相对目录。").String()
	eventTabName := eventTabCmd.Flag("name", "生成的事件表名。").String()
	eventTabDefEventRecursion := eventTabCmd.Flag("default_event_recursion", "生成的事件表代码默认的事件递归处理方式，事件选项[EventRecursion_Allow][EventRecursion_Disallow][EventRecursion_NotEmit][EventRecursion_Discard][EventRecursion_Deepest]可以覆盖此配置。").Default("EventRecursion_Discard").String()

	cmd := kingpin.Parse()

	ctx := &_CommandContext{}
	ctx.DeclFile, _ = filepath.Abs(*declFile)
	ctx.EventRegexp = strings.TrimSpace(*eventRegexp)
	ctx.EventPackageAlias = strings.TrimSpace(*eventPackageAlias)

	if ctx.EventPackageAlias == "" {
		panic("`gen_emit --event_package_alias`设置的别名不能为空")
	}

	switch cmd {
	case emitCmd.FullCommand():
		loadDeclFile(ctx)

		ctx.EmitPackage = strings.TrimSpace(*emitPackage)
		ctx.EmitDir = strings.TrimSpace(*emitDir)
		ctx.EmitDefExport, _ = strconv.ParseBool(*emitDefExport)

		if ctx.EmitPackage == "" {
			panic("`gen_emit --package`设置的包名不能为空")
		}

		genEmit(ctx)
		return
	case eventTabCmd.FullCommand():
		loadDeclFile(ctx)

		ctx.EventTabPackage = *eventTabPackage
		ctx.EventTabDir = *eventTabDir
		ctx.EventTabName = *eventTabName
		ctx.EventTabDefEventRecursion = *eventTabDefEventRecursion

		if ctx.EventTabPackage == "" {
			panic("`gen_eventtab --package`设置的包名不能为空")
		}

		if ctx.EventTabName == "" {
			panic("`gen_eventtab --name`设置的事件表名不能为空")
		}

		if ctx.EventTabDefEventRecursion == "" {
			panic("`gen_eventtab --default_event_recursion`设置的事件递归处理方式不能为空")
		}

		genEventTab(ctx)
		return
	default:
		kingpin.Usage()
		return
	}
}

func loadDeclFile(ctx *_CommandContext) {
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
