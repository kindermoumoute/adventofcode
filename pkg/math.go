package pkg

func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func Check(err error) {
	if err != nil {
		panic(err)
	}
}
