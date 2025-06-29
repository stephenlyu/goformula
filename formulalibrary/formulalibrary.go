package formulalibrary

import (
	"errors"
	"strings"
	"sync"

	"github.com/stephenlyu/goformula/datalibrary"
	. "github.com/stephenlyu/goformula/formulalibrary/base/factory"
	"github.com/stephenlyu/goformula/formulalibrary/base/formula"
	"github.com/stephenlyu/goformula/formulalibrary/native/nativefactory"
	"github.com/stephenlyu/goformula/formulalibrary/native/nativeformulas"
	"github.com/stephenlyu/goformula/stockfunc/function"
	"github.com/z-ray/log"
)

type FormulaChangeListener interface {
	OnFormulaChanged(name string)
}

type FormulaLibrary struct {
	formulaFactories map[string]FormulaCreatorFactory
	lock             sync.Mutex

	debug bool

	loadNative   bool // 是否加载Native公式
	loadLua      bool // 是否加载Lua公式
	loadEasyLang bool // 是否加载EasyLang公式

	formulaChangedListeners []FormulaChangeListener

	dataLibrary datalibrary.DataLibrary
}

func newFormulaLibrary() *FormulaLibrary {
	this := &FormulaLibrary{}
	this.Reset()
	return this
}

func (this *FormulaLibrary) SetDebug(v bool) {
	this.debug = v
}

func (this *FormulaLibrary) SetLoadNative(v bool) {
	this.loadNative = v
}

func (this *FormulaLibrary) SetDataLibrary(dl datalibrary.DataLibrary) {
	this.dataLibrary = dl
	datalibrary.SetDataLibrary(dl)
}

func (this *FormulaLibrary) Reset() {
	this.formulaFactories = make(map[string]FormulaCreatorFactory)
	this.formulaChangedListeners = []FormulaChangeListener{}
	this.debug = false
	this.loadNative = true
	this.loadLua = true
	this.loadEasyLang = true
	this.dataLibrary = nil
}

func (this *FormulaLibrary) CanSupport(name string) bool {
	this.lock.Lock()
	defer this.lock.Unlock()
	_, ok := this.formulaFactories[name]
	return ok
}

func (this *FormulaLibrary) LoadNativeFormulas() {
	if !this.loadNative {
		log.Warnf("Please turn loadNative on first!")
		return
	}

	for _, nf := range nativeformulas.NativeFormulas {
		name := strings.ToUpper(nf.Meta.Name)
		_, factory := nativefactory.NewNativeFormulaCreatorFactory(name)
		this.Register(name, factory)
	}
}

func (this *FormulaLibrary) LoadAllFormulas(luaDir string, easyLangDir string) {
	if this.loadNative {
		this.LoadNativeFormulas()
	}
}

func (this *FormulaLibrary) IsFormulaExisted(name string) bool {
	this.lock.Lock()
	defer this.lock.Unlock()
	_, ok := this.formulaFactories[name]
	return ok
}

func (this *FormulaLibrary) Register(name string, creatorFactory FormulaCreatorFactory) {
	this.lock.Lock()
	defer this.lock.Unlock()
	this.formulaFactories[name] = creatorFactory
}

func (this *FormulaLibrary) Unregister(name string, creatorFactory FormulaCreatorFactory) {
	this.lock.Lock()
	defer this.lock.Unlock()
	delete(this.formulaFactories, name)
}

func (this *FormulaLibrary) GetCreatorFactory(name string) FormulaCreatorFactory {
	this.lock.Lock()
	defer this.lock.Unlock()
	return this.formulaFactories[name]
}

// FormulaChangedListener相关Routine

func (this *FormulaLibrary) AddFormulaChangeListener(listener FormulaChangeListener) {
	if listener == nil {
		return
	}
	this.lock.Lock()
	defer this.lock.Unlock()
	for _, l := range this.formulaChangedListeners {
		if l == listener {
			return
		}
	}
	this.formulaChangedListeners = append(this.formulaChangedListeners, listener)
}

func (this *FormulaLibrary) RemoveFormulaChangeListener(listener FormulaChangeListener) {
	if listener == nil {
		return
	}
	this.lock.Lock()
	defer this.lock.Unlock()
	for i, l := range this.formulaChangedListeners {
		if l == listener {
			this.formulaChangedListeners = append(this.formulaChangedListeners[:i], this.formulaChangedListeners[i+1:]...)
			return
		}
	}
}

func (this *FormulaLibrary) NotifyFormulaChanged(name string) {
	this.lock.Lock()
	defer this.lock.Unlock()
	for _, listener := range this.formulaChangedListeners {
		listener.OnFormulaChanged(name)
	}
}

// 实现FormulaManager接口

func (this *FormulaLibrary) CanSupportVar(name string, varName string) bool {
	this.lock.Lock()
	defer this.lock.Unlock()
	factory, ok := this.formulaFactories[name]
	if !ok {
		return false
	}

	return factory.GetMeta().HasVar(varName)
}

func (this *FormulaLibrary) CanSupportPeriod(period string) bool {
	return true
}

func (this *FormulaLibrary) CanSupportSecurity(code string) bool {
	return true
}

func (this *FormulaLibrary) NewFormula(name string, data function.RVectorReader) formula.Formula {
	factory, ok := this.formulaFactories[name]
	if !ok {
		panic(errors.New("公式不存在"))
	}

	args := factory.GetMeta().DefaultArgs()
	return this.NewFormulaWithArgs(name, data, args)
}

// TODO: 1. 如何缓存公式 2. 如何管理公式间的依赖 3. 公式改变后，如何更新依赖
func (this *FormulaLibrary) NewFormulaWithArgs(name string, data function.RVectorReader, args []float64) formula.Formula {
	factory, ok := this.formulaFactories[name]
	if !ok {
		panic(errors.New("公式不存在"))
	}

	creator := factory.CreateFormulaCreator(args)
	err, formula := creator.CreateFormula(data)
	if err != nil {
		panic(err)
	}
	return formula
}

var GlobalLibrary = newFormulaLibrary()
