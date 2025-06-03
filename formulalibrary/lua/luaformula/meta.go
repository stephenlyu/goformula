package luaformula

import (
	"github.com/aarzilli/golua/lua"
	. "github.com/stephenlyu/goformula/formulalibrary/base/formula"
	"github.com/stevedonovan/luar"
)

func getFormulaDesc(L *lua.State) (name string, argNames []string, args []Arg, flags []int, colors []*Color, lineThick []int, lineStyles []int, graphTypes []int, vars []string) {
	L.GetField(-1, "name")
	luar.LuaToGo(L, -1, &name)
	L.Pop(1)

	L.GetField(-1, "argName")
	luar.LuaToGo(L, -1, &argNames)
	L.Pop(1)

	if len(argNames) > 0 {
		args = make([]Arg, len(argNames))
		var values []float64
		for i, argName := range argNames {
			L.GetField(-1, argName)
			luar.LuaToGo(L, -1, &values)
			L.Pop(1)

			args[i].Default = values[0]
			args[i].Min = values[1]
			args[i].Max = values[2]
		}
	}

	L.GetField(-1, "flags")
	luar.LuaToGo(L, -1, &flags)
	L.Pop(1)

	L.GetField(-1, "color")
	luar.LuaToGo(L, -1, &colors)
	for i, color := range colors {
		if color.Red == -1 {
			colors[i] = nil
		}
	}
	L.Pop(1)

	L.GetField(-1, "lineThick")
	luar.LuaToGo(L, -1, &lineThick)
	L.Pop(1)

	L.GetField(-1, "lineStyle")
	luar.LuaToGo(L, -1, &lineStyles)
	L.Pop(1)

	L.GetField(-1, "graphType")
	luar.LuaToGo(L, -1, &graphTypes)
	L.Pop(1)

	L.GetField(-1, "vars")
	luar.LuaToGo(L, -1, &vars)
	L.Pop(1)

	return
}

func GetMetaFromLuaState(L *lua.State, meta *FormulaMetaImpl) {
	name, argNames, argDefs, flags, colors, lineThick, lineStyles, graphTypes, vars := getFormulaDesc(L)

	meta.Name = name
	meta.ArgNames = argNames
	meta.ArgMeta = argDefs
	meta.Flags = flags
	meta.Colors = colors
	meta.LineThicks = lineThick
	meta.LineStyles = lineStyles
	meta.GraphTypes = graphTypes
	meta.Vars = vars
}
