package easylang

import (
	"regexp"
	"strconv"
	"strings"
	"github.com/stephenlyu/goformula/formulalibrary/base/formula"
)

var ColorLiterals = []string{
	"COLORBLACK",
	"COLORBLUE",
	"COLORGREEN",
	"COLORCYAN",
	"COLORRED",
	"COLORMAGENTA",
	"COLORBROWN",
	"COLORLIGRAY",
	"COLORLIBLUE",
	"COLORLIGREEN",
	"COLORLICYAN",
	"COLORLIRED",
	"COLORLIMAGENTA",
	"COLORYELLOW",
	"COLORWHITE",
}

var colorDefinitions = map[string]*formula.Color{
	"COLORBLACK":     {0, 0, 0},
	"COLORBLUE":      {0, 0, 255},
	"COLORGREEN":     {0, 255, 0},
	"COLORCYAN":      {0, 255, 255},
	"COLORRED":       {255, 0, 0},
	"COLORMAGENTA":   {255, 0, 255},
	"COLORBROWN":     {165, 42, 42},
	"COLORLIGRAY":    {211, 211, 211},
	"COLORLIBLUE":    {173, 216, 230},
	"COLORLIGREEN":   {144, 238, 144},
	"COLORLICYAN":    {224, 255, 255},
	"COLORLIRED":     {255, 0, 128},
	"COLORLIMAGENTA": {255, 128, 128},
	"COLORYELLOW":    {255, 255, 0},
	"COLORWHITE":     {255, 255, 255},
}

var colorRegexp, _ = regexp.Compile("^COLOR([0-F]{2})([0-F]{2})([0-F]{2})$")
var lineThickRegexp, _ = regexp.Compile("^LINETHICK[1-9]$")

func IsValidColorLiteral(color string) bool {
	if colorRegexp.Match([]byte(color)) {
		return true
	}

	for i := range ColorLiterals {
		if ColorLiterals[i] == color {
			return true
		}
	}
	return false
}

func ParseColorLiteral(s string) *formula.Color {
	subMatch := colorRegexp.FindSubmatch([]byte(s))
	if len(subMatch) > 0 {
		b, _ := strconv.Atoi(string(subMatch[1]))
		g, _ := strconv.Atoi(string(subMatch[2]))
		r, _ := strconv.Atoi(string(subMatch[3]))
		return &formula.Color{Red: r, Green: g, Blue: b}
	}

	color := colorDefinitions[s]
	if color == nil {
		panic("Invalid color literal")
	}

	return color
}

func IsValidDescription(desc string) bool {
	switch {
	case desc == "DRAWABOVE":
		return true
	case desc == "NOFRAME":
		return true
	case desc == "NODRAW":
		return true
	case desc == "NOTEXT":
		return true
	case desc == "COLORSTICK":
		return true
	case desc == "STICK":
		return true
	case desc == "LINESTICK":
		return true
	case desc == "VOLSTICK":
		return true
	case desc == "DOTLINE":
		return true
	case desc == "CROSSDOT":
		return true
	case desc == "CIRCLEDOT":
		return true
	case desc == "POINTDOT":
		return true
	case strings.HasPrefix(desc, "COLOR"):
		return IsValidColorLiteral(desc)
	case strings.HasPrefix(desc, "LINETHICK"):
		return lineThickRegexp.Match([]byte(desc))
	}
	return false
}
