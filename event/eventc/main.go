/*
 * This file is part of Golaxy Distributed Service Development Framework.
 *
 * Golaxy Distributed Service Development Framework is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 2.1 of the License, or
 * (at your option) any later version.
 *
 * Golaxy Distributed Service Development Framework is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public License
 * along with Golaxy Distributed Service Development Framework. If not, see <http://www.gnu.org/licenses/>.
 *
 * Copyright (c) 2024 pangdogs.
 */

// Package eventc 使用 go:generate 功能，在编译前自动化生成代码
/*
	- 可以生成事件（Event）与事件表（Event Table）代码。
	- 用于生成事件代码时，在事件定义代码源文件（.go）头部，添加以下注释：
		//go:generate go run git.golaxy.org/core/event/eventc event
	- 用于生成事件表代码时，在事件定义代码源文件（.go）头部，添加以下注释：
		//go:generate go run git.golaxy.org/core/event/eventc eventtab --name={事件表名称}
	- 在 cmd 控制台中，进入事件定义代码源文件（.go）的目录，输入 go generate 指令即可生成代码，此外也可以使用 IDE 提供的 go generate 功能。
	- 编译本包并执行 eventc --help ，可以查看命令行参数，通过参数可以调整生成的代码。
*/
package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func main() {
	rootCmd := &cobra.Command{
		Short: "Event code generation tool.",
		PreRun: func(cmd *cobra.Command, args []string) {
			viper.BindPFlags(cmd.Flags())
			{
				declFile := viper.GetString("decl_file")
				if declFile == "" {
					log.Panic("[--decl_file] must not be empty")
				}
				stat, err := os.Stat(declFile)
				if err != nil {
					log.Panicf("[--decl_file] invalid file: %s", err)
				}
				if stat.IsDir() {
					log.Panic("[--decl_file] must reference a file, not a directory")
				}
			}
			{
				eventRegexp := viper.GetString("event_regexp")
				if eventRegexp == "" {
					log.Panic("[--event_regexp] must not be empty")
				}
			}
			{
				packageEventAlias := viper.GetString("package_event_alias")
				if packageEventAlias == "" {
					log.Panic("[--package_event_alias] must not be empty")
				}
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: true,
		},
	}

	rootCmd.PersistentFlags().String("decl_file", os.Getenv("GOFILE"), "Event declaration source file (.go).")
	rootCmd.PersistentFlags().String("event_regexp", "^[eE]vent.+", "Regular expression used to match event declarations.")
	rootCmd.PersistentFlags().String("package_event_alias", "event", fmt.Sprintf("Import alias used for the Golaxy event package (%s).", packageEventPath))
	rootCmd.PersistentFlags().Bool("copyright", true, "Include the Golaxy Distributed Service Development Framework copyright notice.")

	eventCmd := &cobra.Command{
		Use:   "event",
		Short: "Generate event code from declared events. Supported declaration options: +event-gen:export_emit=[0,1]&auto=[0,1].",
		PreRun: func(cmd *cobra.Command, args []string) {
			viper.BindPFlags(cmd.Flags())
			loadDeclFile()
		},
		Run: func(cmd *cobra.Command, args []string) {
			genEvent()
		},
	}
	eventCmd.Flags().Bool("default_export_emit", true, "Default visibility of generated emit helpers. Can be overridden by +event-gen:export_emit=[0,1].")
	eventCmd.Flags().Bool("default_auto", true, "Generate simplified auto-binding helpers by default. Can be overridden by +event-gen:auto=[0,1].")

	eventTabCmd := &cobra.Command{
		Use:   "eventtab",
		Short: "Generate event table code from declared events. Supported declaration options: +event-tab-gen:recursion=[allow,disallow,discard,skip_received,receive_once].",
		PreRun: func(cmd *cobra.Command, args []string) {
			viper.BindPFlags(cmd.Flags())
			{
				pkg := viper.GetString("package")
				if pkg == "" {
					log.Panic("[eventtab --package] must not be empty")
				}
			}
			{
				dir := viper.GetString("dir")
				if dir == "" {
					log.Panic("[eventtab --dir] must not be empty")
				}
			}
			{
				name := viper.GetString("name")
				if name == "" {
					log.Panic("[eventtab --name] must not be empty")
				}
			}
			loadDeclFile()
		},
		Run: func(cmd *cobra.Command, args []string) {
			genEventTab()
		},
	}
	eventTabCmd.Flags().String("package", os.Getenv("GOPACKAGE"), "Package name used for generated event table code.")
	eventTabCmd.Flags().String("dir", filepath.Dir(os.Getenv("GOFILE")), "Output directory for generated event table source files (.go).")
	eventTabCmd.Flags().String("name", defaultEventTab(), "Name of the generated event table.")
	eventTabCmd.Flags().Bool("export_interface", true, "Control whether the generated event table interface is exported.")

	rootCmd.AddCommand(eventCmd, eventTabCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Panic(err)
	}
}
