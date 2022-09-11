package main

import (
	"fmt"
	"go/ast"
	"regexp"
)

type _EventDecl struct {
	Name               string
	Comment            string
	FuncName           string
	FuncParamsDecl     string
	FuncParams         string
	FuncTypeParamsDecl string
	FuncTypeParams     string
	FuncHasRet         bool
}

type _EventDeclTab []_EventDecl

func (tab *_EventDeclTab) Parse(ctx *_CommandContext) {
	eventRegexp, err := regexp.Compile(ctx.EventRegexp)
	if err != nil {
		panic(err)
	}

	ast.Inspect(ctx.FileAst, func(node ast.Node) bool {
		ts, ok := node.(*ast.TypeSpec)
		if !ok {
			return true
		}

		eventName := ts.Name.Name
		var eventComment string

		for _, comment := range ctx.FileAst.Comments {
			if ctx.FileSet.Position(comment.End()).Line+1 == ctx.FileSet.Position(node.Pos()).Line {
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

				begin := ctx.FileSet.Position(param.Type.Pos())
				end := ctx.FileSet.Position(param.Type.End())

				eventFuncParamsDecl += fmt.Sprintf(", %s %s", paramName, ctx.FileData[begin.Offset:end.Offset])
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

				begin := ctx.FileSet.Position(typeParam.Type.Pos())
				end := ctx.FileSet.Position(typeParam.Type.End())

				if eventFuncTypeParamsDecl != "" {
					eventFuncTypeParamsDecl += ", "
				}
				eventFuncTypeParamsDecl += fmt.Sprintf("%s %s", typeParamName, ctx.FileData[begin.Offset:end.Offset])
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

		eventInfo := _EventDecl{
			Name:               eventName,
			Comment:            eventComment,
			FuncName:           eventFuncName,
			FuncParamsDecl:     eventFuncParamsDecl,
			FuncParams:         eventFuncParams,
			FuncTypeParamsDecl: eventFuncTypeParamsDecl,
			FuncTypeParams:     eventFuncTypeParams,
			FuncHasRet:         eventFuncHasRet,
		}

		*tab = append(*tab, eventInfo)

		return true
	})
}
