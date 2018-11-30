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

func Check(err error) {
	if err != nil {
		panic(err)
	}
}
