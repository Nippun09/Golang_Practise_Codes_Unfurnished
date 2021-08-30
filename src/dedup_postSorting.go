package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")
	//Array = {5,4,1,3,3,2,1}
	//Remove duplicates without using map and sort the result array without using sort function

	input := []int{5, 4, 1, 3, 3, 2, 1} //first sorting and then de-duplicating makes task very easy.
	Quicksort(&input)

	//n := len(input)
	fmt.Println("sorted array", input)

	slint := RemoveDuplicate(input)

	fmt.Println("deduplicated array: ", slint)
}

func RemoveDuplicate(slint []int) []int {

	for i := 0; i < len(slint)-2; i++ {
		if slint[i] == slint[i+1] {
			slint = append(slint[0:i+1], slint[i+2:]...)
		}
	}

	return slint
}

func Quicksort(sliceptr *[]int) {

	index := Pivotize(sliceptr)

	if index != 0 {
		part1 := (*sliceptr)[0:index]
		Quicksort(&part1)
	}

	if index < len(*sliceptr)-1 {
		part2 := (*sliceptr)[index+1:]
		Quicksort(&part2)

	}

}

func Pivotize(sliceptr *[]int) int {

	pivot := (*sliceptr)[0]
	index := 0
	n := len(*sliceptr)
	var swap int

	for i := n - 1; i >= 0; i-- {

		if (*sliceptr)[i] < pivot && index < i {
			swap = (*sliceptr)[index]
			(*sliceptr)[index] = (*sliceptr)[i]
			(*sliceptr)[i] = swap
			index = i
			pivot = (*sliceptr)[index]

		}

		if (*sliceptr)[i] > pivot && i < index {

			swap = (*sliceptr)[index]
			(*sliceptr)[index] = (*sliceptr)[i]
			(*sliceptr)[i] = swap
			index = i
			pivot = (*sliceptr)[index]

		}
	}

	return index

}
