package formulalibrary

import (
	"errors"
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"sync"

	"github.com/stephenlyu/goformula/datalibrary"
	. "github.com/stephenlyu/goformula/formulalibrary/base/factory"
	"github.com/stephenlyu/goformula/formulalibrary/base/formula"
	"github.com/stephenlyu/goformula/formulalibrary/native"
	"github.com/stephenlyu/goformula/formulalibrary/native/nativefactory"
	"github.com/stephenlyu/goformula/stockfunc/function"
	"github.com/stephenlyu/tds/util"
	"github.com/z-ray/log"
)

type FormulaChangeListener interface {
	OnFormulaChanged(name string)
}

type FormulaLibrary[T float32 | float64] struct {
	formulaFactories map[string]FormulaCreatorFactory[T]
	lock             sync.Mutex

	formulaCache map[string]formula.Formula[T]
	idSequence   int64

	nativeFormulas []native.NativeFormula[T]

	debug bool

	loadNative   bool // 是否加载Native公式
	loadLua      bool // 是否加载Lua公式
	loadEasyLang bool // 是否加载EasyLang公式

	formulaChangedListeners []FormulaChangeListener

	dataLibrary datalibrary.DataLibrary[T]
}

func NewFormulaLibrary[T float32 | float64]() *FormulaLibrary[T] {
	this := &FormulaLibrary[T]{}
	this.Reset()
	return this
}

func (this *FormulaLibrary[T]) SetDebug(v bool) {
	this.debug = v
}

func (this *FormulaLibrary[T]) SetNativeFormulas(nativeFormulas []native.NativeFormula[T]) {
	this.nativeFormulas = nativeFormulas
}

func (this *FormulaLibrary[T]) SetLoadNative(v bool) {
	this.loadNative = v
}

func (this *FormulaLibrary[T]) SetDataLibrary(dl datalibrary.DataLibrary[T]) {
	this.dataLibrary = dl
}

func (this *FormulaLibrary[T]) GetDataLibrary() datalibrary.DataLibrary[T] {
	return this.dataLibrary
}

func (this *FormulaLibrary[T]) Reset() {
	this.formulaFactories = make(map[string]FormulaCreatorFactory[T])
	this.formulaChangedListeners = []FormulaChangeListener{}
	this.formulaCache = make(map[string]formula.Formula[T])
	this.debug = false
	this.loadNative = true
	this.loadLua = true
	this.loadEasyLang = true
	this.dataLibrary = nil
}

func (this *FormulaLibrary[T]) CanSupport(name string) bool {
	this.lock.Lock()
	defer this.lock.Unlock()
	_, ok := this.formulaFactories[name]
	return ok
}

func (this *FormulaLibrary[T]) LoadNativeFormulas() {
	if !this.loadNative {
		log.Warnf("Please turn loadNative on first!")
		return
	}

	for i := range this.nativeFormulas {
		nf := &this.nativeFormulas[i]
		name := strings.ToUpper(nf.Meta.Name)
		_, factory := nativefactory.NewNativeFormulaCreatorFactory(this, nf)
		this.Register(name, factory)
	}
}

func (this *FormulaLibrary[T]) LoadAllFormulas(luaDir string, easyLangDir string) {
	if this.loadNative {
		this.LoadNativeFormulas()
	}
}

func (this *FormulaLibrary[T]) IsFormulaExisted(name string) bool {
	this.lock.Lock()
	defer this.lock.Unlock()
	_, ok := this.formulaFactories[name]
	return ok
}

func (this *FormulaLibrary[T]) Register(name string, creatorFactory FormulaCreatorFactory[T]) {
	this.lock.Lock()
	defer this.lock.Unlock()
	this.formulaFactories[name] = creatorFactory
}

func (this *FormulaLibrary[T]) Unregister(name string, creatorFactory FormulaCreatorFactory[T]) {
	this.lock.Lock()
	defer this.lock.Unlock()
	delete(this.formulaFactories, name)
}

func (this *FormulaLibrary[T]) GetCreatorFactory(name string) FormulaCreatorFactory[T] {
	this.lock.Lock()
	defer this.lock.Unlock()
	return this.formulaFactories[name]
}

// FormulaChangedListener相关Routine

func (this *FormulaLibrary[T]) AddFormulaChangeListener(listener FormulaChangeListener) {
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

func (this *FormulaLibrary[T]) RemoveFormulaChangeListener(listener FormulaChangeListener) {
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

func (this *FormulaLibrary[T]) NotifyFormulaChanged(name string) {
	this.lock.Lock()
	defer this.lock.Unlock()
	for _, listener := range this.formulaChangedListeners {
		listener.OnFormulaChanged(name)
	}
}

// 实现FormulaManager接口

func (this *FormulaLibrary[T]) CanSupportVar(name string, varName string) bool {
	this.lock.Lock()
	defer this.lock.Unlock()
	factory, ok := this.formulaFactories[name]
	if !ok {
		return false
	}

	return factory.GetMeta().HasVar(varName)
}

func (this *FormulaLibrary[T]) CanSupportPeriod(period string) bool {
	return true
}

func (this *FormulaLibrary[T]) CanSupportSecurity(code string) bool {
	return true
}

func (this *FormulaLibrary[T]) getCacheKey(name string, data function.RVectorReader[T], args []float64) string {
	va := reflect.ValueOf(data)
	ptr := va.Pointer()
	parts := []string{
		name,
		fmt.Sprintf("%x", ptr),
	}
	for _, arg := range args {
		parts = append(parts, util.FormatFloat64(arg))
	}

	return strings.Join(parts, "-")
}

func (this *FormulaLibrary[T]) cachedFormula(name string, data function.RVectorReader[T], args []float64) formula.Formula[T] {
	key := this.getCacheKey(name, data, args)
	f, ok := this.formulaCache[key]
	if ok {
		log.Infof("use cached formula for %s", name)
		return f
	}
	return nil
}

func (this *FormulaLibrary[T]) cacheFormula(name string, data function.RVectorReader[T], args []float64,
	formula formula.Formula[T]) {
	key := this.getCacheKey(name, data, args)
	this.formulaCache[key] = formula
}

func (this *FormulaLibrary[T]) releaseFormula(f formula.Formula[T]) {
	this.lock.Lock()
	defer this.lock.Unlock()
	var foundKey string
	for key, f1 := range this.formulaCache {
		if f1.Id() == f.Id() {
			foundKey = key
			break
		}
	}
	if foundKey != "" {
		delete(this.formulaCache, foundKey)
	}
}

func (this *FormulaLibrary[T]) getId() int64 {
	this.idSequence++
	return this.idSequence
}

func (this *FormulaLibrary[T]) NewFormula(name string, data function.RVectorReader[T]) formula.Formula[T] {
	factory, ok := this.formulaFactories[name]
	if !ok {
		panic(errors.New("公式不存在"))
	}

	args := factory.GetMeta().DefaultArgs()
	return this.NewFormulaWithArgs(name, data, args)
}

// TODO: 1. 如何缓存公式 2. 如何管理公式间的依赖 3. 公式改变后，如何更新依赖
func (this *FormulaLibrary[T]) NewFormulaWithArgs(name string, data function.RVectorReader[T], args []float64) formula.Formula[T] {
	factory, ok := this.formulaFactories[name]
	if !ok {
		panic(errors.New("公式不存在"))
	}

	f := this.cachedFormula(name, data, args)
	if f != nil {
		f.IncRef()
		return f
	}

	creator := factory.CreateFormulaCreator(args)
	err, formula := creator.CreateFormula(data)
	if err != nil {
		panic(err)
	}
	this.cacheFormula(name, data, args, formula)
	formula.SetId(this.getId())
	formula.IncRef()
	runtime.SetFinalizer(formula, this.releaseFormula)
	return formula
}
