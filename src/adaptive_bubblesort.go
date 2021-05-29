package main

import (
	"fmt"
)

func main() {
	Input := []int{6, 5, 4, 2, 1, 10, 3, 7, 8, 9}
	Bubblesort(&Input)
	fmt.Println("sorted array is :", Input)
}

func Bubblesort(input *[]int) {
	n := len(*input)
	var swap int
	number_of_swaps := 0

	for i := 0; i < n; i++ {
		number_of_swaps = 0
		for j := 0; j < n-i-1; j++ {
			if (*input)[j] > (*input)[j+1] {
				swap = (*input)[j+1]
				(*input)[j+1] = (*input)[j]
				(*input)[j] = swap
				number_of_swaps++
			}

		}

		if number_of_swaps == 0 {
			break
		}
	}

	return
}
