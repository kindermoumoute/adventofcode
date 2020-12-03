package pkg

func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

type nullableInt struct {
	value int
}

func Max(values ...int) int {
	if len(values) == 0 {
		panic("no value in max function")
	}

	var max *nullableInt
	for _, value := range values {
		if max == nil || max.value < value {
			max = &nullableInt{value}
		}
	}
	return max.value
}

func Min(values ...int) int {
	if len(values) == 0 {
		panic("no value in min function")
	}

	var max *nullableInt
	for _, value := range values {
		if max == nil || max.value > value {
			max = &nullableInt{value}
		}
	}
	return max.value
}

func Sum(values ...int) int {
	if len(values) == 0 {
		panic("no value in sum function")
	}

	sum := 0
	for _, value := range values {
		sum += value
	}
	return sum
}

func Multiply(values ...int) int {
	if len(values) == 0 {
		panic("no value in sum function")
	}

	sum := 1
	for _, value := range values {
		sum *= value
	}
	return sum
}

func GCD(a, b int) int {
	if b == 0 {
		return a
	}
	return GCD(b, a%b)
}

func LCM(a, b int) int {
	return a / GCD(a, b) * b
}
