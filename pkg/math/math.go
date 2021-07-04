package math

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Max(values ...int) int {
	var max int
	for i, v := range values {
		if i == 0 || v > max {
			max = v
		}
	}
	return max
}

func Min(values ...int) int {
	var min int
	for i, v := range values {
		if i == 0 || v < min {
			min = v
		}
	}
	return min
}

func Pow(base, n int) int {
	ret := 1
	for n > 0 {
		if n&1 == 1 {
			ret *= base
		}
		base *= base
		n >>= 1
	}
	return ret
}

func GCD(x, y int) int {
	if y == 0 {
		return x
	}
	return GCD(y, x%y)
}
