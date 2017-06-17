package function


func sum(values Value, start, end int) float64 {
	sum := 0.0
	for i := start; i < end; i++ {
		sum += values.Get(i)
	}
	return sum
}

func ma_(values Value, start, end int) float64 {
	if start >= end {
		return 0
	}

	return sum(values, start, end) / float64(end - start)
}
