package nativeformulas

import (
	. "github.com/stephenlyu/goformula/formulalibrary/base/formula"
	"github.com/stephenlyu/goformula/function"
	stockfunc "github.com/stephenlyu/goformula/stockfunc/function"
	"github.com/stephenlyu/tds/period"
)

type BaseNativeFormula[T float32 | float64] struct {
	*FormulaMetaImpl
	RefValues__   []function.Value[T]
	DrawActions__ []DrawAction
	Data__        stockfunc.RVectorReader[T]
	RefCount      int64
	id            int64

	formula_index_map map[Formula[T]]*stockfunc.IndexMap[T]
}

func (this *BaseNativeFormula[T]) SetFormulaIndexMap(formula Formula[T], indexMap *stockfunc.IndexMap[T]) {
	// TODO: 非线程安全
	if this.formula_index_map == nil {
		this.formula_index_map = make(map[Formula[T]]*stockfunc.IndexMap[T])
	}
	this.formula_index_map[formula] = indexMap
}

func (this *BaseNativeFormula[T]) GetFormulaIndexMap(formula Formula[T]) *stockfunc.IndexMap[T] {
	if this.formula_index_map == nil {
		return nil
	}
	return this.formula_index_map[formula]
}

func (this *BaseNativeFormula[T]) IncRef() {
	this.RefCount++
}

func (this *BaseNativeFormula[T]) DecRef() {
	this.RefCount--
}

func (this *BaseNativeFormula[T]) SetId(id int64) {
	this.id = id
}

func (this *BaseNativeFormula[T]) Id() int64 {
	return this.id
}

func (this *BaseNativeFormula[T]) Period() period.Period {
	return this.Data__.Period()
}

func (this *BaseNativeFormula[T]) Len() int {
	return this.Data__.Len()
}

func (this BaseNativeFormula[T]) Get(index int) []T {
	ret := make([]T, len(this.RefValues__))
	for i, v := range this.RefValues__ {
		ret[i] = v.Get(index)
	}

	return ret
}

func (this BaseNativeFormula[T]) Ref(offset int) []T {
	index := this.Data__.Len() - 1 - offset
	return this.Get(index)
}

func (this BaseNativeFormula[T]) GetVarValue(varName string) function.Value[T] {
	for i, v := range this.FormulaMetaImpl.Vars {
		if varName == v {
			return this.RefValues__[i]
		}
	}
	return nil
}

func (this *BaseNativeFormula[T]) Destroy() {
}

func (this *BaseNativeFormula[T]) Name() string {
	return this.GetName()
}

func (this *BaseNativeFormula[T]) DrawActions() []DrawAction {
	return this.DrawActions__
}

func (this *BaseNativeFormula[T]) DumpState() {
}
