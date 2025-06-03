package function

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
	"time"

	"github.com/stephenlyu/tds/util"
)

func TestBARSLAST(t *testing.T) {
	values := make([]float64, 10000)
	//rand.Seed(0)
	for i := range values {
		if rand.Int()%30 == 0 {
			values[i] = 1
		}
	}
	v := Vector(values)

	s := time.Now().UnixNano()
	b1 := BARSLASTOLD(v)
	t1 := time.Now().UnixNano() - s
	s = time.Now().UnixNano()
	b2 := BARSLAST(v)
	t2 := time.Now().UnixNano() - s

	fmt.Printf("%dns %dns %d\n", t1, t2, t1/t2)

	for i := range values {
		//fmt.Println(values[i], b1.Get(i), b2.Get(i))
		util.Assert((math.IsNaN(b1.Get(i)) && math.IsNaN(b2.Get(i))) || b1.Get(i) == b2.Get(i), "")
	}
}

func TestBARSLASTS(t *testing.T) {
	values := make([]float64, 200)
	rand.Seed(0)
	for i := range values {
		if rand.Int()%30 == 0 {
			values[i] = 1
		}
	}
	v := Vector(values)
	b := BARSLASTS(v, Scalar(2))

	for i := range values {
		fmt.Println(i, values[i], b.Get(i), i-int(b.Get(i)))
	}

	v.Append(1)
	b.UpdateLastValue()
	for i := range v.Values {
		fmt.Println(i, v.Get(i), b.Get(i), i-int(b.Get(i)))
	}
}
