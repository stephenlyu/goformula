package test

import (
	. "github.com/onsi/ginkgo"
	"io/ioutil"
	"encoding/json"
	stockfunc "github.com/stephenlyu/goformula/stockfunc/function"
	"fmt"
	"time"
	"github.com/stephenlyu/goformula/stockfunc/function"
	"github.com/stephenlyu/goformula/formulalibrary/native/nativefactory"
	"github.com/stephenlyu/tds/date"
	_ "github.com/stephenlyu/goformula/test/native"
)

type Record struct {
	Date string				`json:"date"`
	Open float32			`json:"open"`
	Close float32			`json:"close"`
	High float32			`json:"high"`
	Low float32				`json:"low"`
	Volume float32			`json:"volume"`
	Amount float32			`json:"amount"`
}

func (this *Record) GetUTCDate() uint64 {
	ret, _ := date.DayString2Timestamp(this.Date)
	return ret
}

func (this *Record) GetDate() string {
	return this.Date
}

func (this *Record) GetOpen() float32 {
	return this.Open
}

func (this *Record) GetClose() float32 {
	return this.Close
}

func (this *Record) GetHigh() float32 {
	return this.High
}

func (this *Record) GetLow() float32 {
	return this.Low
}

func (this *Record) GetAmount() float32 {
	return this.Amount
}

func (this *Record) GetVolume() float32 {
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
	It("test", func (){
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
		fmt.Println("time cost: ", (time.Now().UnixNano() - start) / 1000000, "ms")
	})
})
