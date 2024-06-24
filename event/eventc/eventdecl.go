package main

import (
	"fmt"
	"github.com/spf13/viper"
	"go/ast"
	"go/token"
	"regexp"
)

type EventDecl struct {
	Name               string
	Comment            string
	FuncName           string
	FuncParamsDecl     string
	FuncParams         string
	FuncTypeParamsDecl string
	FuncTypeParams     string
	FuncHasRet         bool
}

type EventDeclTab struct {
	Package string
	Events  []EventDecl
}

func (tab *EventDeclTab) Parse() {
	eventRegexp, err := regexp.Compile(viper.GetString("event_regexp"))
	if err != nil {
		panic(err)
	}

	fast := viper.Get("file_ast").(*ast.File)
	fset := viper.Get("file_set").(*token.FileSet)
	fdata := viper.Get("file_data").([]byte)

	tab.Package = fast.Name.Name

	ast.Inspect(fast, func(node ast.Node) bool {
		ts, ok := node.(*ast.TypeSpec)
		if !ok {
			return true
		}

		eventName := ts.Name.Name
		var eventComment string

		for _, comment := range fast.Comments {
			if fset.Position(comment.End()).Line+1 == fset.Position(node.Pos()).Line {
				eventComment = comment.Text()
				break
			}
		}

		if !eventRegexp.MatchString(eventName) {
			return true
		}

		eventIface, ok := ts.Type.(*ast.InterfaceType)
		if !ok {
			return true
		}

		if eventIface.Methods.NumFields() <= 0 {
			return true
		}

		eventFuncField := eventIface.Methods.List[0]

		if len(eventFuncField.Names) <= 0 {
			return true
		}

		eventFuncName := eventFuncField.Names[0].Name

		eventFunc, ok := eventFuncField.Type.(*ast.FuncType)
		if !ok {
			return true
		}

		eventFuncParamsDecl := ""
		eventFuncParams := ""

		if eventFunc.Params != nil {
			for i, param := range eventFunc.Params.List {
				paramName := ""

				for _, pn := range param.Names {
					if paramName != "" {
						paramName += ", "
					}
					paramName += pn.Name
				}

				if paramName == "" {
					paramName = fmt.Sprintf("p%d", i)
				}

				if eventFuncParams != "" {
					eventFuncParams += ", "
				}
				eventFuncParams += paramName

				begin := fset.Position(param.Type.Pos()).Offset
				end := fset.Position(param.Type.End()).Offset

				eventFuncParamsDecl += fmt.Sprintf(", %s %s", paramName, fdata[begin:end])
			}
		}

		eventFuncTypeParamsDecl := ""
		eventFuncTypeParams := ""

		if ts.TypeParams != nil {
			for i, typeParam := range ts.TypeParams.List {
				typeParamName := ""

				for _, pn := range typeParam.Names {
					if typeParamName != "" {
						typeParamName += ", "
					}
					typeParamName += pn.Name
				}

				if typeParamName == "" {
					typeParamName = fmt.Sprintf("p%d", i)
				}

				if eventFuncTypeParams != "" {
					eventFuncTypeParams += ", "
				}
				eventFuncTypeParams += typeParamName

				begin := fset.Position(typeParam.Type.Pos()).Offset
				end := fset.Position(typeParam.Type.End()).Offset

				if eventFuncTypeParamsDecl != "" {
					eventFuncTypeParamsDecl += ", "
				}
				eventFuncTypeParamsDecl += fmt.Sprintf("%s %s", typeParamName, fdata[begin:end])
			}
		}

		if eventFuncTypeParamsDecl != "" {
			eventFuncTypeParamsDecl = fmt.Sprintf("[%s]", eventFuncTypeParamsDecl)
		}

		if eventFuncTypeParams != "" {
			eventFuncTypeParams = fmt.Sprintf("[%s]", eventFuncTypeParams)
		}

		eventFuncHasRet := false

		if eventFunc.Results.NumFields() > 0 {
			eventRet, ok := eventFunc.Results.List[0].Type.(*ast.Ident)
			if !ok {
				return true
			}

			if eventRet.Name != "bool" {
				return true
			}

			eventFuncHasRet = true
		}

		eventDecl := EventDecl{
			Name:               eventName,
			Comment:            eventComment,
			FuncName:           eventFuncName,
			FuncParamsDecl:     eventFuncParamsDecl,
			FuncParams:         eventFuncParams,
			FuncTypeParamsDecl: eventFuncTypeParamsDecl,
			FuncTypeParams:     eventFuncTypeParams,
			FuncHasRet:         eventFuncHasRet,
		}

		tab.Events = append(tab.Events, eventDecl)

		return true
	})
}
