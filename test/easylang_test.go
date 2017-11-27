package test
import (
	. "github.com/onsi/ginkgo"
	"github.com/stephenlyu/goformula/easylang"
	"github.com/stephenlyu/goformula/stockfunc/function"
	"fmt"
	"time"
	"github.com/stephenlyu/goformula/formulalibrary/easylang/easylangfactory"
	"github.com/stephenlyu/goformula/formulalibrary/base/formula"
	"github.com/stephenlyu/goformula/formulalibrary"
	"github.com/stephenlyu/tds/period"
)

var _ = Describe("Compile", func() {
	It("test", func () {
		err := easylang.Compile("MACDBUY.d", "output.lua", nil, true, true)
		if err != nil {
			fmt.Println(err)
		}
	})
})

var _ = Describe("CrossValue", func() {
	It("test", func () {
		err := easylang.Compile("CROSS.d", "cross.lua", nil, true, true)
		if err != nil {
			fmt.Println(err)
		}
	})
})

var _ = Describe("Token", func() {
	It("test", func () {
		err := easylang.Tokenizer("MACD.d")
		if err != nil {
			fmt.Println(err)
		}
	})
})

var _ = Describe("EasyLangMACD", func() {
	It("test", func () {
		_, data := loadJson("300666.SZ.json")
		rv := function.RecordVector(data)

		formulas := []string {"MACD.d", "MA.d", "VOL.d"}
		for _, name := range formulas {
			fmt.Println("Test formula", name, "...")
			start := time.Now().UnixNano()

			err, factory := easylangfactory.NewEasyLangFormulaCreatorFactory(name, nil, true)
			if err != nil {
				panic(err)
			}
			creator := factory.CreateFormulaCreator(nil)

			err, formula := creator.CreateFormula(rv)
			if err != nil {
				panic(err)
			}
			defer formula.Destroy()

			for i := 0; i < formula.VarCount(); i++ {
				fmt.Printf("name: %s noDraw: %v lineThick: %d color: %+v\n", formula.VarName(i), formula.NoDraw(i), formula.LineThick(i), formula.Color(i))
			}

			start1 := time.Now().UnixNano()
			len := formula.Len()
			fmt.Println("formula.len:", len, "data len:", rv.Len())
			for i := 0; i < len; i++ {
				r := formula.Get(i)
				fmt.Printf("%s %+v\n", rv.Get(i).GetDate(), r)
			}

			fmt.Println("time cost: ", (time.Now().UnixNano() - start) / 1000000, (time.Now().UnixNano() - start1) / 1000000, "ms")
		}
	})
})

var _ = Describe("ELMACDBuy", func() {
	It("test", func (){
		_, data := loadJson("300666.SZ.json")
		_, p := period.PeriodFromString("D1")
		rv := function.RecordVectorEx("300666", p, data)

		fmt.Println("data len:", len(data))
		start := time.Now().UnixNano()

		library := formulalibrary.GlobalLibrary
		library.Reset()
		library.LoadEasyLangFormulas(".")

		f := library.NewFormula("MACDBUY", rv)
		defer f.Destroy()

		fmt.Println("Name:", f.GetName())
		for i := 0; i < f.ArgCount(); i++ {
			min, max := f.ArgRange(i)
			fmt.Printf("default: %f min: %f max: %f\n", f.ArgDefault(i), min, max)
		}
		for i := 0; i < f.VarCount(); i++ {
			fmt.Printf("name: %s noDraw: %v lineThick: %d color: %+v\n", f.VarName(i), f.NoDraw(i), f.LineThick(i), f.Color(i))
		}

		fmt.Println(f.Len())
		for i := 0; i < f.Len(); i++ {
			r := f.Get(i)
			fmt.Printf("%d. %s\t%.02f\n", i, rv.Get(i).GetDate(), r[0])
		}
		fmt.Println("time cost: ", (time.Now().UnixNano() - start) / 1000000, "ms")
	})
})

var _ = Describe("ELDDGS", func() {
	It("test", func (){
		_, data := loadJson("300666.SZ.json")
		_, p := period.PeriodFromString("D1")
		rv := function.RecordVectorEx("300666", p, data)

		fmt.Println("data len:", rv.Len())
		start := time.Now().UnixNano()

		library := formulalibrary.GlobalLibrary
		library.Reset()
		library.SetDebug(true)
		library.LoadEasyLangFormulas(".")

		f := library.NewFormula("DDGS", rv)
		defer f.Destroy()

		fmt.Println("Name:", f.GetName())
		for i := 0; i < f.ArgCount(); i++ {
			min, max := f.ArgRange(i)
			fmt.Printf("default: %f min: %f max: %f\n", f.ArgDefault(i), min, max)
		}
		for i := 0; i < f.VarCount(); i++ {
			fmt.Printf("name: %s noDraw: %v lineThick: %d color: %+v\n", f.VarName(i), f.NoDraw(i), f.LineThick(i), f.Color(i))
		}

		fmt.Println(f.Len())
		for i := 0; i < f.Len(); i++ {
			r := f.Get(i)
			fmt.Printf("%d. %s\t%.02f\n", i, rv.Get(i).GetDate(), r[0])
		}
		fmt.Println("time cost: ", (time.Now().UnixNano() - start) / 1000000, "ms")
	})
})

var _ = Describe("ELDrawLine", func() {
	It("test", func (){
		_, data := loadJson("300666.SZ.json")
		rv := function.RecordVector(data)

		fmt.Println("data len:", len(data))
		start := time.Now().UnixNano()

		err, factory := easylangfactory.NewEasyLangFormulaCreatorFactory("DRAWLINE.d", nil, true)
		if err != nil {
			panic(err)
		}
		creator := factory.CreateFormulaCreator(nil)

		err, f := creator.CreateFormula(rv)
		if err != nil {
			panic(err)
		}
		defer f.Destroy()

		fmt.Println("Name:", f.GetName())
		for i := 0; i < f.ArgCount(); i++ {
			min, max := f.ArgRange(i)
			fmt.Printf("default: %f min: %f max: %f\n", f.ArgDefault(i), min, max)
		}
		for i := 0; i < f.VarCount(); i++ {
			fmt.Printf("name: %s noDraw: %v lineThick: %d color: %+v\n", f.VarName(i), f.NoDraw(i), f.LineThick(i), f.Color(i))
		}

		for i := 0; i < len(f.DrawActions()); i++ {
			a := f.DrawActions()[i].(formula.DrawLine)
			fmt.Println(a.GetActionType(), a.GetColor(), a.GetLineThick(), a.GetVarIndex())
		}

		fmt.Println(f.Len())
		for i := 0; i < f.Len(); i++ {
			r := f.Get(i)
			fmt.Printf("%d. %s\t%.02f\t%.02f\t%.02f\t%.02f\t%.02f\t%.02f\t%.02f\n", i, rv.Get(i).GetDate(), r[0], r[1], r[2], r[3], r[4], r[5], r[6])
		}
		fmt.Println("time cost: ", (time.Now().UnixNano() - start) / 1000000, "ms")
	})
})

var _ = Describe("ELPloyLine", func() {
	It("test", func (){
		_, data := loadJson("300666.SZ.json")
		rv := function.RecordVector(data)

		fmt.Println("data len:", len(data))
		start := time.Now().UnixNano()

		err, factory := easylangfactory.NewEasyLangFormulaCreatorFactory("PLOYLINE.d", nil, true)
		if err != nil {
			panic(err)
		}
		creator := factory.CreateFormulaCreator(nil)

		err, f := creator.CreateFormula(rv)
		if err != nil {
			panic(err)
		}
		defer f.Destroy()

		fmt.Println("Name:", f.GetName())
		for i := 0; i < f.ArgCount(); i++ {
			min, max := f.ArgRange(i)
			fmt.Printf("default: %f min: %f max: %f\n", f.ArgDefault(i), min, max)
		}
		for i := 0; i < f.VarCount(); i++ {
			fmt.Printf("name: %s noDraw: %v lineThick: %d color: %+v\n", f.VarName(i), f.NoDraw(i), f.LineThick(i), f.Color(i))
		}

		fmt.Println(f.Len())
		for i := 0; i < f.Len(); i++ {
			r := f.Get(i)
			fmt.Printf("%d. %s\t%.02f\n", i, rv.Get(i).GetDate(), r[0])
		}
		fmt.Println("time cost: ", (time.Now().UnixNano() - start) / 1000000, "ms")
	})
})

var _ = Describe("ELDrawActions", func() {
	It("test", func (){
		_, data := loadJson("300666.SZ.json")
		rv := function.RecordVector(data)

		fmt.Println("data len:", len(data))
		start := time.Now().UnixNano()

		err, factory := easylangfactory.NewEasyLangFormulaCreatorFactory("DRAW.d", nil, true)
		if err != nil {
			panic(err)
		}
		creator := factory.CreateFormulaCreator(nil)

		err, f := creator.CreateFormula(rv)
		if err != nil {
			panic(err)
		}
		defer f.Destroy()

		fmt.Println("Name:", f.GetName())
		for i := 0; i < f.ArgCount(); i++ {
			min, max := f.ArgRange(i)
			fmt.Printf("default: %f min: %f max: %f\n", f.ArgDefault(i), min, max)
		}
		for i := 0; i < f.VarCount(); i++ {
			fmt.Printf("name: %s noDraw: %v lineThick: %d color: %+v\n", f.VarName(i), f.NoDraw(i), f.LineThick(i), f.Color(i))
		}

		drawActions := f.DrawActions()
		for _, action := range drawActions {
			fmt.Printf("%+v\n", action)
		}

		fmt.Println("time cost: ", (time.Now().UnixNano() - start) / 1000000, "ms")
	})
})
