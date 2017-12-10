package main

import "fmt"

func main() {
	puzzle := []int{197, 97, 204, 108, 1, 29, 5, 71, 0, 50, 2, 255, 248, 78, 254, 63}
	puzzle2 := []int{'1', '9', '7', ',', '9', '7', ',', '2', '0', '4', ',', '1', '0', '8', ',', '1', ',', '2', '9', ',', '5', ',', '7', '1', ',', '0', ',', '5', '0', ',', '2', ',', '2', '5', '5', ',', '2', '4', '8', ',', '7', '8', ',', '2', '5', '4', ',', '6', '3', 17, 31, 73, 47, 23}
	list := newList(256)
	knotHash(list, puzzle, 0, 0)

	fmt.Println(list[0] * list[1])

	list = newList(256)
	pos, skipSize := 0, 0
	for i := 0; i < 64; i++ {
		pos, skipSize = knotHash(list, puzzle2, pos, skipSize)
	}
	hash := make([]int, 16)
	for i, v := range list {
		hash[i/len(hash)] ^= v
	}
	for _, h := range hash {
		if h < 16 {
			fmt.Printf("0")
		}
		fmt.Printf("%x", h)
	}

}

func newList(ln int) []int {
	list := make([]int, ln)
	for i := range list {
		list[i] = i
	}
	return list
}

func knotHash(list, puzzle []int, pos, skipSize int) (int, int) {

	//fmt.Println("New hasing...")
	for _, length := range puzzle {
		//fmt.Println("length ", length)
		for i := 0; i < length/2; i++ {
			swapStart := (pos + i) % len(list)
			swapEnd := (pos + length - i - 1) % len(list)
			//fmt.Println("end:", swapEnd, " start:", swapStart, " pos:", pos, " i:", i)
			list[swapStart], list[swapEnd] = list[swapEnd], list[swapStart]
		}
		pos = (pos + length + skipSize) % len(list)
		skipSize = (skipSize + 1) % len(list)
	}
	return pos, skipSize
}
