package test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	. "github.com/onsi/ginkgo"
	"github.com/stephenlyu/goformula/formulalibrary/native/nativefactory"
	"github.com/stephenlyu/goformula/stockfunc/function"
	stockfunc "github.com/stephenlyu/goformula/stockfunc/function"
	"github.com/stephenlyu/tds/date"
)

type Record struct {
	Date   string  `json:"date"`
	Open   float64 `json:"open"`
	Close  float64 `json:"close"`
	High   float64 `json:"high"`
	Low    float64 `json:"low"`
	Volume float64 `json:"volume"`
	Amount float64 `json:"amount"`
}

func (this *Record) GetUTCDate() uint64 {
	ret, _ := date.DayString2Timestamp(this.Date)
	return ret
}

func (this *Record) GetDate() string {
	return this.Date
}

func (this *Record) GetOpen() float64 {
	return this.Open
}

func (this *Record) GetClose() float64 {
	return this.Close
}

func (this *Record) GetHigh() float64 {
	return this.High
}

func (this *Record) GetLow() float64 {
	return this.Low
}

func (this *Record) GetAmount() float64 {
	return this.Amount
}

func (this *Record) GetVolume() float64 {
	return this.Volume
}

func loadJson(jsonFile string) (error, []function.Record) {
	bytes, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		return err, nil
	}

	var result []Record
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return err, nil
	}

	ret := make([]function.Record, len(result))

	for i := range result {
		ret[i] = &result[i]
	}

	return err, ret
}

var _ = Describe("NMACD", func() {
	It("test", func() {
		_, data := loadJson("data.json")
		rv := stockfunc.RecordVector(data)

		start := time.Now().UnixNano()

		_, factory := nativefactory.NewNativeFormulaCreatorFactory("MACD")
		creator := factory.CreateFormulaCreator(nil)
		_, macd := creator.CreateFormula(rv)

		fmt.Println("Name:", macd.GetName())
		for i := 0; i < macd.ArgCount(); i++ {
			min, max := macd.ArgRange(i)
			fmt.Printf("default: %f min: %f max: %f\n", macd.ArgDefault(i), min, max)
		}
		for i := 0; i < macd.VarCount(); i++ {
			fmt.Printf("name: %s noDraw: %v lineThick: %d color: %+v\n", macd.VarName(i), macd.NoDraw(i), macd.LineThick(i), macd.Color(i))
		}

		for i := 0; i < macd.Len(); i++ {
			r := macd.Get(i)
			fmt.Printf("%s\t%.02f\t%.02f\t%.02f\n", rv.Get(i).GetDate(), r[0], r[1], r[2])
		}
		fmt.Println("time cost: ", (time.Now().UnixNano()-start)/1000000, "ms")
	})
})
