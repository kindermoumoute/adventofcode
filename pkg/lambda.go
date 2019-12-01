package pkg

type SumFunc func(int) (int, bool)

func RecursiveSum(i int, sumFunc SumFunc) int {
	var stop bool
	i, stop = sumFunc(i)
	if stop {
		return 0
	}
	return i + RecursiveSum(i, sumFunc)
}
