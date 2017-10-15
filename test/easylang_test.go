package test
import (
	. "github.com/onsi/ginkgo"
	"github.com/stephenlyu/goformula/easylang"
	"github.com/stephenlyu/goformula/stockfunc/function"
	"fmt"
	"time"
	"github.com/stephenlyu/goformula/stockfunc"
)

var _ = Describe("Compile", func() {
	It("test", func () {
		err := easylang.Compile("MACD.d", "output.lua")
		if err != nil {
			fmt.Println(err)
		}
	})
})

var _ = Describe("Token", func() {
	It("test", func () {
		err := easylang.Tokenizer("MACD.d")
		if err != nil {
			fmt.Println(err)
		}
	})
})

var _ = Describe("EasyLangMACD", func() {
	It("test", func () {
		_, data := loadJson("data.json")
		rv := function.RecordVector(data["300666"])

		factory := stockfunc.NewFormulaFactory(true)

		formulas := []string {"MACD.d", "MA.d", "VOL.d"}
		args := [][]float64{
			[]float64{12, 26, 9},
			[]float64{5, 10, 20, 60},
			[]float64{5, 10},
		}

		for i, name := range formulas {
			fmt.Println("Test formula", name, "...")
			start := time.Now().UnixNano()

			err, formula := factory.NewEasyLangFormula(name, rv, args[i])
			if err != nil {
				panic(err)
			}
			defer formula.Destroy()
			for i := 0; i < formula.VarCount(); i++ {
				fmt.Printf("name: %s noDraw: %v lineThick: %d color: %+v\n", formula.VarName(i), formula.NoDraw(i), formula.LineThick(i), formula.Color(i))
			}

			start1 := time.Now().UnixNano()
			len := formula.Len()
			fmt.Println("formula.len:", len, "data len:", rv.Len())
			for i := 0; i < len; i++ {
				r := formula.Get(i)
				fmt.Printf("%s %+v\n", rv.Get(i).GetDate(), r)
			}

			fmt.Println("time cost: ", (time.Now().UnixNano() - start) / 1000000, (time.Now().UnixNano() - start1) / 1000000, "ms")
		}
	})
})
