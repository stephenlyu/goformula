package test
import (
	. "github.com/onsi/ginkgo"
	"github.com/stephenlyu/goformula/stockfunc/function"
	"fmt"
	"time"
	"github.com/stephenlyu/goformula/formulalibrary"
)

var _ = Describe("LuaMACD", func() {
	It("test", func () {
		_, data := loadJson("data.json")
		rv := function.RecordVector(data)

		start := time.Now().UnixNano()

		library := formulalibrary.GlobalLibrary
		library.Reset()
		library.LoadLuaFormulas("luas")

		formula := library.NewFormula("MACD", rv)
		defer formula.Destroy()

		fmt.Println("Name:", formula.Name())
		for i := 0; i < formula.ArgCount(); i++ {
			min, max := formula.ArgRange(i)
			fmt.Printf("default: %f min: %f max: %f\n", formula.ArgDefault(i), min, max)
		}
		for i := 0; i < formula.VarCount(); i++ {
			fmt.Printf("name: %s noDraw: %v lineThick: %d color: %+v\n", formula.VarName(i), formula.NoDraw(i), formula.LineThick(i), formula.Color(i))
		}

		start1 := time.Now().UnixNano()
		len := formula.Len()
		fmt.Println("macd.len:", len, "data len:", rv.Len())
		for i := 0; i < len; i++ {
			r := formula.Get(i)
			fmt.Printf("%s\t%.02f\t%.02f\t%.02f\t%.02f\t%.02f\n", rv.Get(i).GetDate(), r[0], r[1], r[2], r[3], r[4])
		}

		fmt.Println("time cost: ", (time.Now().UnixNano() - start) / 1000000, (time.Now().UnixNano() - start1) / 1000000, "ms")
	})
})

var _ = Describe("LuaDrawLine", func() {
	It("test", func () {
		_, data := loadJson("300666.SZ.json")
		rv := function.RecordVector(data)

		start := time.Now().UnixNano()

		library := formulalibrary.GlobalLibrary
		library.Reset()
		library.LoadLuaFormulas("luas")

		formula := library.NewFormula("DRAWLINE", rv)
		defer formula.Destroy()

		fmt.Println("Name:", formula.Name())
		for i := 0; i < formula.ArgCount(); i++ {
			min, max := formula.ArgRange(i)
			fmt.Printf("default: %f min: %f max: %f\n", formula.ArgDefault(i), min, max)
		}
		for i := 0; i < formula.VarCount(); i++ {
			fmt.Printf("name: %s noDraw: %v lineThick: %d color: %+v\n", formula.VarName(i), formula.NoDraw(i), formula.LineThick(i), formula.Color(i))
		}

		start1 := time.Now().UnixNano()
		len := formula.Len()
		fmt.Println("formula.len:", len, "data len:", rv.Len())
		for i := 0; i < len; i++ {
			r := formula.Get(i)
			fmt.Printf("%d. %s\t%.02f\n", i, rv.Get(i).GetDate(), r[0])
		}

		fmt.Println("time cost: ", (time.Now().UnixNano() - start) / 1000000, (time.Now().UnixNano() - start1) / 1000000, "ms")
	})
})
