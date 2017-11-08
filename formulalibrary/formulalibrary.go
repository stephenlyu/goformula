package formulalibrary

import (
	. "github.com/stephenlyu/goformula/formulalibrary/base/factory"
	"github.com/stephenlyu/goformula/formulalibrary/easylang/easylangfactory"
)

type FormulaLibrary struct {
	formulas map[string]FormulaCreatorFactory
}

func newFormulaLibrary() *FormulaLibrary {
	return &FormulaLibrary {
		formulas: make(map[string]FormulaCreatorFactory),
	}
}

func (this *FormulaLibrary) RegisterEasyLangFile(name, easyLangFile string, debug bool) error {
	err, factory := easylangfactory.NewEasyLangFormulaCreatorFactory(easyLangFile, debug)
	if err != nil {
		return err
	}
	this.Register(name, factory)
	return nil
}

func (this *FormulaLibrary) Register(name string, creatorFactory FormulaCreatorFactory) {
	this.formulas[name] = creatorFactory
}

func (this *FormulaLibrary) Unregister(name string, creatorFactory FormulaCreatorFactory) {
	delete(this.formulas, name)
}

func (this *FormulaLibrary) GetCreatorFactory(name string) FormulaCreatorFactory {
	return this.formulas[name]
}

func (this *FormulaLibrary) CanSupport(name string) bool {
	_, ok := this.formulas[name]
	return ok
}

var GlobalLibrary = newFormulaLibrary()
