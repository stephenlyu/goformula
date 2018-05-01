package function

import (
	"testing"
	"math/rand"
	"github.com/stephenlyu/tds/util"
	"time"
	"fmt"
	"math"
)

func TestBARSLAST(t *testing.T) {
	values := make([]float64, 10000)
	//rand.Seed(0)
	for i := range values {
		if rand.Int() % 30  == 0 {
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

	fmt.Printf("%dns %dns %d\n", t1, t2, t1 / t2)

	for i := range values {
		//fmt.Println(values[i], b1.Get(i), b2.Get(i))
		util.Assert((math.IsNaN(b1.Get(i)) && math.IsNaN(b2.Get(i))) || b1.Get(i) == b2.Get(i), "")
	}
}
