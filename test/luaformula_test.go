package test
import (
	. "github.com/onsi/ginkgo"
	stockfunc "github.com/stephenlyu/goformula/stockfunc/function"
	"github.com/stephenlyu/goformula/luafunc"
	"fmt"
	"time"
)

var _ = Describe("LuaMACD", func() {
	It("test", func () {
		_, data := loadJson("data.json")
		rv := stockfunc.RecordVector(data["300666"])

		start := time.Now().UnixNano()

		err, formula := luafunc.NewFormula("macd.lua", rv, []float64{12, 26, 9})
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
