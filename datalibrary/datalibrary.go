package datalibrary

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/stephenlyu/goformula/stockfunc/function"
	"github.com/stephenlyu/tds/datasource"
	tdxdatasource "github.com/stephenlyu/tds/datasource/tdx"
	"github.com/stephenlyu/tds/entity"
	. "github.com/stephenlyu/tds/period"
	"github.com/stephenlyu/tds/util"
	"github.com/z-ray/log"
)

const (
	DATA_TYPE_ORIGINAL = iota
	DATA_TYPE_FORWARD_ADJUST
)

type DataLibrary interface {
	// Set data type, original or forward adjust?
	SetDataType(dataType int)

	// Get data with specific code & period
	GetData(code string, period string) function.RVectorReader

	// Release data
	ReleaseData(data function.RVectorReader)
}

type dataLibrary struct {
	dataType int

	dataDir string
	ds      datasource.DataSource
}

func NewDataLibrary(dataDir string) DataLibrary {
	ds := tdxdatasource.NewDataSource(dataDir, true)
	return &dataLibrary{dataDir: dataDir, ds: ds}
}

func (this *dataLibrary) SetDataType(dataType int) {
	util.Assert(dataType == DATA_TYPE_ORIGINAL || dataType == DATA_TYPE_FORWARD_ADJUST, "")
	this.dataType = dataType
}

func (this *dataLibrary) translateCode(code string) string {
	if strings.Index(code, ".") != -1 {
		return code
	}

	code = strings.ToUpper(code)
	if code[0] == '0' || code[0] == '3' {
		return code + ".SZ"
	} else if unicode.IsNumber(rune(code[0])) {
		return code + ".SH"
	}
	if strings.HasPrefix(code, "SZ") {
		return code[2:] + ".SZ"
	}
	if strings.HasPrefix(code, "SH") {
		return code[2:] + ".SH"
	}

	return code + ".SH"
}

func (this *dataLibrary) GetData(code string, periodString string) function.RVectorReader {
	code = this.translateCode(code)
	fmt.Printf("load data: %s %s\n", code, periodString)

	security, err := entity.ParseSecurity(code)
	util.Assert(err == nil, "bad security code")

	err, period := PeriodFromString(periodString)
	util.Assert(err == nil, "bad period")

	var data []entity.Record
	switch this.dataType {
	case DATA_TYPE_ORIGINAL:
		err, data = this.ds.GetData(security, period)
	case DATA_TYPE_FORWARD_ADJUST:
		err, data = this.ds.GetForwardAdjustedData(security, period)
	}

	ret := make([]function.Record, len(data))
	if err != nil {
		log.Errorf("DataLibrary.GetData fail, error: %v", err)
		return function.RecordVectorEx(code, period, ret)
	}

	for i := range data {
		ret[i] = &data[i]
	}

	fmt.Println(data[0].GetDate(), data[len(data)-1].GetDate())

	return function.RecordVectorEx(code, period, ret)
}

func (this *dataLibrary) ReleaseData(data function.RVectorReader) {
}

var GlobalDataLibrary DataLibrary

func SetDataLibrary(dl DataLibrary) {
	GlobalDataLibrary = dl
}
