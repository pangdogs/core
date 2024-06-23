// Package eventc 使用go:generate功能，在编译前自动化生成代码
/*
	- 可以生成事件（event）与事件表（event table）代码。
	- 用于生成事件代码时，在事件定义代码源文件（.go）头部，添加以下注释：
		//go:generate go run git.golaxy.org/core/event/eventc event
	- 用于生成事件表代码时，在事件定义代码源文件（.go）头部，添加以下注释：
		//go:generate go run git.golaxy.org/core/event/eventc eventtab --name={事件表名称}
	- 在cmd控制台中，进入事件定义代码源文件（.go）的目录，输入go generate指令即可生成代码，此外也可以使用IDE提供的go generate功能。
	- 编译本包并执行eventcode --help，可以查看命令行参数，通过参数可以调整生成的代码。
*/
package main

import (
	"fmt"
	"github.com/alecthomas/kingpin/v2"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	// 基础选项
	declFile := kingpin.Flag("decl_file", "定义事件的源文件（.go）。").Default(os.Getenv("GOFILE")).ExistingFile()
	eventRegexp := kingpin.Flag("event_regexp", "匹配事件定义时，使用的正则表达式。").Default("^[eE]vent.+").String()
	packageEventAlias := kingpin.Flag("package_event_alias", fmt.Sprintf("导入Golaxy框架的`%s`包时使用的别名。", packageEventPath)).Default("event").String()
	packageIfaceAlias := kingpin.Flag("package_iface_alias", fmt.Sprintf("导入Golaxy框架的`%s`包时使用的别名。", packageIfacePath)).Default("iface").String()

	// 生成事件代码相关选项
	eventCmd := kingpin.Command("event", "通过定义的事件，生成事件代码。(定义选项：+event-gen:export=[0,1]&auto=[0,1])")
	eventPackage := eventCmd.Flag("package", "生成事件代码时，使用的包名。").Default(os.Getenv("GOPACKAGE")).String()
	eventDir := eventCmd.Flag("dir", "生成事件代码时，输出的源文件（.go）存放的相对目录。").String()
	eventDefExport := eventCmd.Flag("default_export", "生成事件代码时，发送事件的代码的可见性，事件定义选项+event-gen:export=[0,1]可以覆盖此配置。").Default("true").String()
	eventDefAuto := eventCmd.Flag("default_auto", "生成事件代码时，是否生成简化绑定事件的代码，事件定义选项+event-gen:auto=[0,1]可以覆盖此配置。").Default("true").String()

	// 生成事件表代码相关选项
	eventTabCmd := kingpin.Command("eventtab", "通过定义的事件，生成事件表代码。(定义选项：+event-tab-gen:recursion=[allow,disallow,discard,truncate,deepest]")
	eventTabPackage := eventTabCmd.Flag("package", "生成事件表代码，使用的包名。").Default(os.Getenv("GOPACKAGE")).String()
	eventTabDir := eventTabCmd.Flag("dir", "生成事件表代码时，输出的源文件（.go）存放的相对目录。").String()
	eventTabName := eventTabCmd.Flag("name", "生成的事件表名称。").String()

	cmd := kingpin.Parse()

	ctx := &CommandContext{}
	ctx.DeclFile, _ = filepath.Abs(*declFile)
	ctx.EventRegexp = strings.TrimSpace(*eventRegexp)
	ctx.PackageEventAlias = strings.TrimSpace(*packageEventAlias)
	ctx.PackageIfaceAlias = strings.TrimSpace(*packageIfaceAlias)

	if ctx.PackageEventAlias == "" {
		panic("`--package_event_alias`设置的别名不能为空")
	}

	if ctx.PackageIfaceAlias == "" {
		panic("`--package_iface_alias`设置的别名不能为空")
	}

	switch cmd {
	case eventCmd.FullCommand():
		loadDeclFile(ctx)

		ctx.EventPackage = strings.TrimSpace(*eventPackage)
		ctx.EventDir = strings.TrimSpace(*eventDir)
		ctx.EventDefExport, _ = strconv.ParseBool(*eventDefExport)
		ctx.EventDefAuto, _ = strconv.ParseBool(*eventDefAuto)

		if ctx.EventPackage == "" {
			panic("`event --package`设置的包名不能为空")
		}

		genEvent(ctx)
		return
	case eventTabCmd.FullCommand():
		loadDeclFile(ctx)

		ctx.EventTabPackage = strings.TrimSpace(*eventTabPackage)
		ctx.EventTabDir = strings.TrimSpace(*eventTabDir)
		ctx.EventTabName = strings.TrimSpace(*eventTabName)

		if ctx.EventTabPackage == "" {
			panic("`eventtab --package`设置的包名不能为空")
		}

		if ctx.EventTabName == "" {
			panic("`eventtab --name`设置的事件表名不能为空")
		}

		genEventTab(ctx)
		return
	default:
		kingpin.Usage()
		return
	}
}
