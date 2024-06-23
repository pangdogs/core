package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"net/url"
	"path/filepath"
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
	packageEventPath = "git.golaxy.org/core/event"
	packageIfacePath = "git.golaxy.org/core/utils/iface"
)

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

func parseGenAtti(str, atti string) url.Values {
	idx := strings.Index(str, atti)
	if idx < 0 {
		return url.Values{}
	}

	str = str[idx+len(atti):]

	end := strings.IndexAny(str, "\r\n")
	if end >= 0 {
		str = str[:end]
	}

	values, _ := url.ParseQuery(str)

	for k, vs := range values {
		for i, v := range vs {
			vs[i] = strings.TrimSpace(v)
		}
		values[k] = vs
	}

	return values
}
