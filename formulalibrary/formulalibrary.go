package formulalibrary

import (
	. "github.com/stephenlyu/goformula/formulalibrary/base/factory"
	"github.com/stephenlyu/goformula/formulalibrary/easylang/easylangfactory"
	"github.com/stephenlyu/goformula/formulalibrary/base/formula"
	"github.com/stephenlyu/goformula/stockfunc/function"
	"sync"
	"errors"
	"github.com/stephenlyu/goformula/formulalibrary/lua/luaformula"
	"github.com/stephenlyu/goformula/formulalibrary/lua/luafactory"
	"github.com/z-ray/log"
	"github.com/stephenlyu/goformula/formulalibrary/native/nativeformulas"
	"github.com/stephenlyu/goformula/formulalibrary/native/nativefactory"
	"strings"
	"path/filepath"
	"fmt"
	"github.com/stephenlyu/goformula/datalibrary"
)

type FormulaChangeListener interface {
	OnFormulaChanged(name string)
}

type FormulaLibrary struct {
	formulaFactories        map[string]FormulaCreatorFactory
	lock                    sync.Mutex

	debug 					bool

	loadNative 				bool 						// 是否加载Native公式
	loadLua					bool						// 是否加载Lua公式
	loadEasyLang 			bool 						// 是否加载EasyLang公式

	formulaChangedListeners []FormulaChangeListener

	dataLibrary 			datalibrary.DataLibrary
}

func newFormulaLibrary() *FormulaLibrary {
	this := &FormulaLibrary {}
	this.Reset()

	luaformula.SetFormulaManager(this)
	return this
}

func (this *FormulaLibrary) SetDebug(v bool) {
	this.debug = v
}

func (this *FormulaLibrary) SetLoadNative(v bool) {
	this.loadNative = v
}

func (this *FormulaLibrary) SetLoadLua(v bool) {
	this.loadLua = v
}

func (this *FormulaLibrary) SetLoadEasyLang(v bool) {
	this.loadEasyLang = v
}

func (this *FormulaLibrary) SetDataLibrary(dl datalibrary.DataLibrary) {
	this.dataLibrary = dl
	luaformula.SetDataLibrary(dl)
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

func (this *FormulaLibrary) LoadLuaFormulas(dir string) {
	if !this.loadLua {
		log.Warnf("Please turn loadLua on first!")
		return
	}

	files, err := filepath.Glob(filepath.Join(dir, "*.lua"))
	if err != nil {
		log.Errorf("Load lua formulas fail, error: %v", err)
		return
	}

	for _, filePath := range files {
		baseName := filepath.Base(filePath)
		parts := strings.Split(baseName, ".")
		name := strings.ToUpper(parts[0])
		if this.IsFormulaExisted(name) {
			log.Errorf("Formula %s exist, lua version ignored.", name)
			continue
		}

		err = this.RegisterLuaFile(name, filePath)
		if err != nil {
			log.Errorf("Load lua formula %s fail, error: %v", filePath, err)
		} else {
			log.Infof("Load lua formula %s success", filePath)
		}
	}
}

func (this *FormulaLibrary) LoadEasyLangFormulas(dir string) {
	if !this.loadLua {
		log.Warnf("Please turn loadEasyLang on first!")
		return
	}

	files, err := filepath.Glob(filepath.Join(dir, "*.d"))
	if err != nil {
		log.Errorf("Load lua formulas fail, error: %v", err)
		return
	}

	var errorFiles []string

	for {
		changed := false
		errorFiles = []string{}
		for _, filePath := range files {
			baseName := filepath.Base(filePath)
			parts := strings.Split(baseName, ".")
			name := strings.ToUpper(parts[0])
			if this.IsFormulaExisted(name) {
				log.Errorf("Formula %s exist, easylang version ignored.", name)
				continue
			}
			err = this.RegisterEasyLangFile(name, filePath, this.debug)
			if err != nil {
				log.Errorf("Load easy lang formula %s fail, error: %+v", filePath, err)
				errorFiles = append(errorFiles, filePath)
			} else {
				changed = true
			}
		}
		if changed {
			files = errorFiles
		} else {
			break
		}
	}
}

func (this *FormulaLibrary) LoadAllFormulas(luaDir string, easyLangDir string) {
	if this.loadNative {
		this.LoadNativeFormulas()
	}

	if this.loadLua {
		this.LoadLuaFormulas(luaDir)
	}

	if this.loadEasyLang {
		this.LoadEasyLangFormulas(easyLangDir)
	}
}

func (this *FormulaLibrary) RegisterEasyLangFile(name, easyLangFile string, debug bool) (err error) {
	defer func() {
		if err1 := recover(); err1 != nil {
			err = errors.New(fmt.Sprintf("%v", err1))
		}
	}()

	err, factory := easylangfactory.NewEasyLangFormulaCreatorFactory(easyLangFile, this, debug)
	if err != nil {
		return err
	}
	this.Register(name, factory)
	return nil
}

func (this *FormulaLibrary) RegisterLuaFile(name, luaFile string) (err error) {
	defer func() {
		if err1 := recover(); err1 != nil {
			err = errors.New(fmt.Sprintf("%v", err1))
		}
	}()

	err, factory := luafactory.NewLuaFormulaCreatorFactory(luaFile)
	if err != nil {
		return err
	}
	this.Register(name, factory)
	return nil
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

func (this *FormulaLibrary) NewFormula(name string, data *function.RVector) formula.Formula {
	factory, ok := this.formulaFactories[name]
	if !ok {
		panic(errors.New("公式不存在"))
	}

	args := factory.GetMeta().DefaultArgs()
	return this.NewFormulaWithArgs(name, data, args)
}

// TODO: 1. 如何缓存公式 2. 如何管理公式间的依赖 3. 公式改变后，如何更新依赖
func (this *FormulaLibrary) NewFormulaWithArgs(name string, data *function.RVector, args []float64) formula.Formula {
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
