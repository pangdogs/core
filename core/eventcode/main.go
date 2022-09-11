// Package eventcode 事件代码生成器，用于`go:generate`生成事件（Event）相关代码。
//
// - 在事件定义代码源文件（*.go）头部，添加以下注释：
//	go:generate go run github.com/pangdogs/galaxy/core/eventcode --decl_file=$GOFILE gen_emit --package=$GOPACKAGE
//
// - 在事件定义代码源文件（*.go）的路径下，打开cmd控制台运行
//	go generate
// 或使用IDE提供的go generate功能生成代码。
//
// - 可以用于生成事件发送代码与事件表代码。
//
// - 可以编译并运行
//	eventcode --help
// 查看命令行参数，并按自己的需求修改生成的代码。
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
	DeclFile      string
	EventRegexp   string
	CoreAlias     string
	NotImportCore bool
	FileData      []byte
	FileSet       *token.FileSet
	FileAst       *ast.File

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
	coreAlias := kingpin.Flag("core_alias", "导入银河（Galaxy）框架的`github.com/pangdogs/galaxy/core`包时使用的别名。").Default("core").String()
	notImportCore := kingpin.Flag("not_import_core", "是否不导入银河（Galaxy）框架的`github.com/pangdogs/galaxy/core`包。").Default("false").Bool()

	// 生成发送事件代码相关选项
	emitCmd := kingpin.Command("gen_emit", "通过定义的事件生成发送事件代码。")
	emitPackage := emitCmd.Flag("package", "生成的发送事件代码时使用的包名。").String()
	emitDir := emitCmd.Flag("dir", "生成的发送事件代码产生的源文件（*.go）存放的相对目录。").String()
	emitDefExport := emitCmd.Flag("default_export", "生成的发送事件代码默认的可见性。").Default("true").String()

	// 生成事件表代码相关选项
	eventTabCmd := kingpin.Command("gen_eventtab", "通过定义的事件生成事件表代码。")
	eventTabPackage := eventTabCmd.Flag("package", "生成的事件表代码使用的包名。").String()
	eventTabDir := eventTabCmd.Flag("dir", "生成的事件表产生的源文件（*.go）存放的相对目录。").String()
	eventTabName := eventTabCmd.Flag("name", "生成的事件表名。").String()
	eventTabDefEventRecursion := eventTabCmd.Flag("default_event_recursion", "生成的事件表代码默认的事件递归处理方式。").Default("EventRecursion_Discard").String()

	cmd := kingpin.Parse()

	ctx := &_CommandContext{}
	ctx.DeclFile, _ = filepath.Abs(*declFile)
	ctx.EventRegexp = strings.TrimSpace(*eventRegexp)
	ctx.CoreAlias = strings.TrimSpace(*coreAlias)
	ctx.NotImportCore = *notImportCore

	if ctx.CoreAlias == "" {
		panic("`gen_emit --core_alias`设置的别名不能为空")
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
