package test

import (
	. "github.com/onsi/ginkgo"
	//"github.com/stephenlyu/goformular/function"
	//"fmt"
	"io/ioutil"
	"encoding/json"
	stockfunc "github.com/stephenlyu/goformular/stockfunc/function"
	"fmt"
	"github.com/stephenlyu/goformular/stockfunc/formula"
)

func loadJson(jsonFile string) (error, map[string][]stockfunc.Record) {
	bytes, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		return err, nil
	}

	var result map[string][]stockfunc.Record

	err = json.Unmarshal(bytes, &result)
	return err, result
}

var _ = Describe("MACD", func() {
	It("test", func (){
		_, data := loadJson("data.json")
		rv := stockfunc.RecordVector(data["300666"])

		var macd formula.Formula = formula.MACD(rv, nil, nil, nil)
		for i := 0; i < macd.Len(); i++ {
			r := macd.Get(i)
			fmt.Printf("%s\t%.02f\t%.02f\t%.02f\t%.02f\t%.02f\n", rv.Get(i).Date, r[0], r[1], r[2], r[3], r[4])
		}
	})
})
