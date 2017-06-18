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

var _ = Describe("MA", func() {
	It("test", func (){
		data := make([]float64, 100)
		for i := 0; i < 100; i++ {
			data[i] = float64(i)
		}
		a := function.Vector(data)
		result := function.MA(a, nil)
		for i := 0; i < result.Len(); i++ {
			fmt.Println(result.Get(i))
		}

		fmt.Println("============================")

		a.Append(10)
		result.UpdateLastValue()
		for i := 0; i < result.Len(); i++ {
			fmt.Println(result.Get(i))
		}
	})
})

var _ = Describe("LLVBARS", func() {
	It("test", func (){
		data := make([]float64, 100)
		for i := 0; i < 100; i++ {
			data[i] = float64(i)
		}
		a := function.Vector(data)
		result := function.LLVBARS(a, function.Scalar(10))
		for i := 0; i < result.Len(); i++ {
			fmt.Println(result.Get(i))
		}

		fmt.Println("============================")

		a.Append(10)
		result.UpdateLastValue()
		for i := 0; i < result.Len(); i++ {
			fmt.Println(result.Get(i))
		}
	})
})

var _ = Describe("HHVBARS", func() {
	It("test", func (){
		data := make([]float64, 100)
		for i := 0; i < 100; i++ {
			data[i] = float64(i)
		}
		a := function.Vector(data)
		result := function.HHVBARS(a, function.Scalar(10))
		for i := 0; i < result.Len(); i++ {
			fmt.Println(result.Get(i))
		}

		fmt.Println("============================")

		a.Append(10)
		result.UpdateLastValue()
		for i := 0; i < result.Len(); i++ {
			fmt.Println(result.Get(i))
		}
	})
})


var _ = Describe("ROUND2", func() {
	It("test", func (){
		data := []float64{0.4949999, 0.500001, -0.4949999, -0.50000001}
		a := function.Vector(data)
		result := function.ROUND2(a, function.Scalar(2))
		for i := 0; i < result.Len(); i++ {
			fmt.Println(result.Get(i))
		}

		fmt.Println("============================")

		a.Append(3.1415926)
		result.UpdateLastValue()
		for i := 0; i < result.Len(); i++ {
			fmt.Println(result.Get(i))
		}
	})
})


var _ = Describe("ROUND2", func() {
	It("test", func (){
		data := []float64{0.4949999, 0.500001, -0.4949999, -0.50000001}
		a := function.Vector(data)
		result := function.ROUND2(a, function.Scalar(2))
		for i := 0; i < result.Len(); i++ {
			fmt.Println(result.Get(i))
		}

		fmt.Println("============================")

		a.Append(3.1415926)
		result.UpdateLastValue()
		for i := 0; i < result.Len(); i++ {
			fmt.Println(result.Get(i))
		}
	})
})


var _ = Describe("SLOPE", func() {
	It("test", func (){
		data := make([]float64, 100)
		for i := 0; i < 100; i++ {
			data[i] = float64(i) / 2
		}
		a := function.Vector(data)
		result := function.SLOPE(a, function.Scalar(5))
		for i := 0; i < result.Len(); i++ {
			fmt.Println(result.Get(i))
		}

		fmt.Println("============================")

		a.Append(10)
		result.UpdateLastValue()
		for i := 0; i < result.Len(); i++ {
			fmt.Println(result.Get(i))
		}
	})
})
