package nativeformulas

import (
	. "github.com/stephenlyu/goformula/formulalibrary/native"
	"github.com/stephenlyu/goformula/formulalibrary/base/formula"
	"sync"
	"fmt"
	"github.com/z-ray/log"
)

var lock sync.Mutex
var NativeFormulas = []NativeFormula {
}

func RegisterNativeFormula(create CreateFunc, meta *formula.FormulaMetaImpl) error {
	log.Infof("RegisterNativeFormula name: %s", meta.GetName())
	lock.Lock()
	defer lock.Unlock()

	for i := range NativeFormulas {
		if NativeFormulas[i].Meta.GetName() == meta.GetName() {
			return fmt.Errorf("native formula %s already exist", meta.GetName())
		}
	}

	NativeFormulas = append(NativeFormulas, NativeFormula{Creator: create, Meta: meta})
	return nil
}
