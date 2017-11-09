package nativefactory

import (
	"github.com/stephenlyu/goformula/formulalibrary/base/factory"
	"github.com/stephenlyu/goformula/formulalibrary/native"
	"github.com/stephenlyu/goformula/formulalibrary/native/nativeformulas"
	"errors"
	"fmt"
)

type nativeFormulaCreatorFactory struct {
	nativeFormula *native.NativeFormula
}

func NewNativeFormulaCreatorFactory(name string) (error, factory.FormulaCreatorFactory) {
	var nativeFormula *native.NativeFormula
	for i := range nativeformulas.NativeFormulas {
		nf := &nativeformulas.NativeFormulas[i]
		if nf.Meta.Name == name {
			nativeFormula = nf
			break
		}
	}

	if nativeFormula == nil {
		return errors.New(fmt.Sprintf("No %s formula", name)), nil
	}

	return nil, &nativeFormulaCreatorFactory{nativeFormula: nativeFormula}
}
func (this *nativeFormulaCreatorFactory) GetDefaultArgs() []float64 {
	return this.nativeFormula.Meta.DefaultArgs()
}

func (this *nativeFormulaCreatorFactory) CreateFormulaCreator(args []float64) factory.FormulaCreator {
	if args == nil {
		args = this.nativeFormula.Meta.DefaultArgs()
	}

	return &nativeFormulaCreator{
		factory: this,
		args: args,
	}
}
