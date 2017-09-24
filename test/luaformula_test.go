package test
import (
	. "github.com/onsi/ginkgo"
	"github.com/stephenlyu/goformula/stockfunc/function"
	"fmt"
	"time"
	"github.com/stephenlyu/goformula/stockfunc"
)

var _ = Describe("LuaMACD", func() {
	It("test", func () {
		_, data := loadJson("data.json")
		rv := function.RecordVector(data["300666"])

		start := time.Now().UnixNano()

		factory := stockfunc.NewFormulaFactory(false)

		err, formula := factory.NewLuaFormula("macd.lua", rv, []float64{12, 26, 9})
		if err != nil {
			panic(err)
		}
		defer formula.Destroy()

		fmt.Println("Name:", formula.Name())
		for i := 0; i < formula.ArgCount(); i++ {
			min, max := formula.ArgRange(i)
			fmt.Printf("default: %f min: %f max: %f\n", formula.ArgDefault(i), min, max)
		}
		for i := 0; i < formula.VarCount(); i++ {
			fmt.Printf("name: %s noDraw: %d lineThick: %d color: %s\n", formula.VarName(i), formula.NoDraw(i), formula.LineThick(i), formula.Color(i))
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
