package nativeformulas

import (
	"fmt"
	"strings"
	"sync"

	"github.com/stephenlyu/goformula/formulalibrary/base/formula"
	. "github.com/stephenlyu/goformula/formulalibrary/native"
	"github.com/z-ray/log"
)

var lock sync.Mutex
var F64NativeFormulas = []NativeFormula[float64]{}
var F32NativeFormulas = []NativeFormula[float32]{}

func RegisterF64NativeFormula(create CreateFunc[float64], meta *formula.FormulaMetaImpl) error {
	log.Infof("RegisterF64NativeFormula name: %s", meta.GetName())
	lock.Lock()
	defer lock.Unlock()

	for i := range F64NativeFormulas {
		if F64NativeFormulas[i].Meta.GetName() == meta.GetName() {
			return fmt.Errorf("native formula %s already exist", meta.GetName())
		}
	}

	F64NativeFormulas = append(F64NativeFormulas, NativeFormula[float64]{Creator: create, Meta: meta})
	return nil
}

func FindF64NativeFormula(name string) *NativeFormula[float64] {
	for i := range F64NativeFormulas {
		nf := &F64NativeFormulas[i]
		if strings.ToUpper(nf.Meta.Name) == name {
			return nf
		}
	}
	return nil
}

func RegisterF32NativeFormula(create CreateFunc[float32], meta *formula.FormulaMetaImpl) error {
	log.Infof("RegisterF32NativeFormula name: %s", meta.GetName())
	lock.Lock()
	defer lock.Unlock()

	for i := range F32NativeFormulas {
		if F32NativeFormulas[i].Meta.GetName() == meta.GetName() {
			return fmt.Errorf("native formula %s already exist", meta.GetName())
		}
	}

	F32NativeFormulas = append(F32NativeFormulas, NativeFormula[float32]{Creator: create, Meta: meta})
	return nil
}

func FindF32NativeFormula(name string) *NativeFormula[float32] {
	for i := range F32NativeFormulas {
		nf := &F32NativeFormulas[i]
		if strings.ToUpper(nf.Meta.Name) == name {
			return nf
		}
	}
	return nil
}
