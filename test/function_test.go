package test

import (
. "github.com/onsi/ginkgo"
	"github.com/stephenlyu/goformular/function"
	"fmt"
)

var _ = Describe("Not", func() {
	It("test", func (){
		data := []float64{0, 1, 2, 3}
		vector := function.Vector(data)
		result := function.NOT(vector)
		for i := 0; i < result.Len(); i++ {
			fmt.Println(result.Get(i))
		}

		vector.Append(0)
		result.UpdateLastValue()
		for i := 0; i < result.Len(); i++ {
			fmt.Println(result.Get(i))
		}
	})
})

var _ = Describe("Add", func() {
	It("test", func (){
		data := []float64{0, 1, 2, 3}
		data1 := []float64{0, 1, 2, 3}
		a := function.Vector(data)
		b := function.Vector(data1)
		result := function.ADD(a, b)
		for i := 0; i < result.Len(); i++ {
			fmt.Println(result.Get(i))
		}

		a.Append(10)
		b.Append(5)
		result.UpdateLastValue()
		for i := 0; i < result.Len(); i++ {
			fmt.Println(result.Get(i))
		}
	})
})
