package gotils
func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func AbsInt64(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

func AbsFloat64(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

func MaxInt(x, y int) int{
	if x > y{
		return x
	}
	return y
}

func MaxInt64(x, y int64) int64{
	if x > y{
		return x
	}
	return y
}

func MaxFloat64(x, y float64) float64{
	if x > y{
		return x
	}
	return y
}

func MinInt(x, y int) int{
	if x < y{
		return x
	}
	return y
}

func MinInt64(x, y int64) int64{
	if x < y{
		return x
	}
	return y
}

func MinFloat64(x, y float64) float64{
	if x < y{
		return x
	}
	return y
}
