package formula

type Formula interface {
	Len() int
	UpdateLastValue()
	Get(index int) []float64
	Ref(offset int) []float64

	Destroy()
}
