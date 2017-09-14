package test

import (
	. "github.com/onsi/ginkgo"
	"io/ioutil"
	"encoding/json"
	stockfunc "github.com/stephenlyu/goformula/stockfunc/function"
	"fmt"
	"github.com/stephenlyu/goformula/stockfunc/formula"
	"time"
	"github.com/stephenlyu/goformula/stockfunc/function"
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


func loadJson(jsonFile string) (error, map[string][]function.Record) {
	bytes, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		return err, nil
	}

	var result map[string][]Record
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return err, nil
	}

	ret := make(map[string][]function.Record)

	for code, records := range result {
		objs := make([]function.Record, len(records))
		for i := range records {
			objs[i] = &records[i]
		}
		ret[code] = objs
	}

	return err, ret
}

var _ = Describe("MACD", func() {
	It("test", func (){
		_, data := loadJson("data.json")
		rv := stockfunc.RecordVector(data["300666"])

		start := time.Now().UnixNano()

		var macd formula.Formula = formula.MACD(rv, nil, nil, nil)
		for i := 0; i < macd.Len(); i++ {
			r := macd.Get(i)
			fmt.Printf("%s\t%.02f\t%.02f\t%.02f\t%.02f\t%.02f\n", rv.Get(i).GetDate(), r[0], r[1], r[2], r[3], r[4])
		}
		fmt.Println("time cost: ", (time.Now().UnixNano() - start) / 1000000, "ms")
	})
})
