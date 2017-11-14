package datalibrary

import (
	. "github.com/stephenlyu/tds/period"
	"github.com/stephenlyu/goformula/stockfunc/function"
	"baiwenbao.com/arbitrage/util"
	"github.com/stephenlyu/tds/datasource"
	"github.com/stephenlyu/tds/datasource/tdx"
	"github.com/stephenlyu/tds/entity"
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
	GetData(code string, period Period) *function.RVector

	// Release data
	ReleaseData(data *function.RVector)
}

type dataLibrary struct {
	dataType int

	dataDir string
	ds datasource.DataSource
}

func NewDataLibrary(dataDir string) DataLibrary {
	ds := tdxdatasource.NewDataSource(dataDir, true)
	return &dataLibrary{dataDir: dataDir, ds: ds}
}

func (this *dataLibrary) SetDataType(dataType int) {
	util.Assert(dataType == DATA_TYPE_ORIGINAL || dataType == DATA_TYPE_FORWARD_ADJUST, "")
	this.dataType = dataType
}

func (this *dataLibrary) GetData(code string, period Period) *function.RVector {
	security, err := entity.ParseSecurity(code)
	util.Assert(err == nil, "bad security code")

	var data []entity.Record
	switch this.dataType {
	case DATA_TYPE_ORIGINAL:
		err, data = this.ds.GetData(security, period)
	case DATA_TYPE_FORWARD_ADJUST:
		err, data = this.ds.GetForwardAdjustedData(security, period)
	}

	if err != nil {
		log.Errorf("DataLibrary.GetData fail, error: %v", err)
	}

	ret := make([]function.Record, len(data))
	for i := range data {
		ret[i] = &data[i]
	}

	return function.RecordVector(ret)
}

func (this *dataLibrary) ReleaseData(data *function.RVector) {
}
