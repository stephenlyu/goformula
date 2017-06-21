package test
import (
	. "github.com/onsi/ginkgo"
	"github.com/stephenlyu/goformula/easylang"
	stockfunc "github.com/stephenlyu/goformula/stockfunc/function"
	"fmt"
	"time"
	"github.com/stephenlyu/goformula/luafunc"
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
		rv := stockfunc.RecordVector(data["300666"])

		err := easylang.Compile("MACD.d", "output.lua")
		if err != nil {
			panic(err)
		}

		start := time.Now().UnixNano()

		err, formula := luafunc.NewFormula("output.lua", rv, []float64{12, 26, 9})
		if err != nil {
			panic(err)
		}
		defer formula.Destroy()

		start1 := time.Now().UnixNano()
		len := formula.Len()
		fmt.Println("macd.len:", len, "data len:", rv.Len())
		for i := 0; i < len; i++ {
			r := formula.Get(i)
			fmt.Printf("%s\t%.02f\t%.02f\t%.02f\t%.02f\t%.02f\n", rv.Get(i).Date, r[0], r[1], r[2], r[3], r[4])
		}

		fmt.Println("time cost: ", (time.Now().UnixNano() - start) / 1000000, (time.Now().UnixNano() - start1) / 1000000, "ms")
	})
})
