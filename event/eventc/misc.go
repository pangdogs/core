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

package main

import (
	"bytes"
	"github.com/spf13/viper"
	"go/parser"
	"go/token"
	"io/ioutil"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

const (
	packageEventPath = "git.golaxy.org/core/event"
	packageIfacePath = "git.golaxy.org/core/utils/iface"
)

const (
	copyrightNotice = `/*
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

`
)

func loadDeclFile() {
	declFile := viper.GetString("decl_file")

	fileData, err := ioutil.ReadFile(declFile)
	if err != nil {
		panic(err)
	}

	fset := token.NewFileSet()

	fast, err := parser.ParseFile(fset, declFile, fileData, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	viper.Set("file_data", fileData)
	viper.Set("file_set", fset)
	viper.Set("file_ast", fast)
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

func snake2Camel(s string) string {
	var buf bytes.Buffer
	upper := true
	for _, c := range s {
		if c == '_' {
			upper = true
			continue
		}
		if upper {
			buf.WriteRune(unicode.ToUpper(c))
			upper = false
		} else {
			buf.WriteRune(c)
		}
	}
	return buf.String()
}

func truncateDot(s string) string {
	idx := strings.Index(s, ".")
	if idx < 0 {
		return s
	}
	return s[:idx]
}

func defaultEventTab() string {
	s := strings.TrimSuffix(strings.TrimSuffix(truncateDot(snake2Camel(filepath.Base(os.Getenv("GOFILE")))), "Event"), "EventTab")
	return s + "EventTab"
}
