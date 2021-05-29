package main

import (
	"fmt"
)

func main() {
	Input := []int{6, 5, 4, 2, 1, 10, 3, 7, 8, 9}
	Quicksort(&Input)
	fmt.Println("sorted array is :", Input)
}

func Quicksort(input *[]int) {
	if len(*input) <= 1 {
		return
	}

	sortedindex := Partition(input)

	first := (*input)[0:sortedindex]
	second := (*input)[sortedindex+1:]

	if len(first) > 0 {
		Quicksort(&first)
	}
	if len(second) > 0 {
		Quicksort(&second)
	}

}

func Partition(sliceptr *[]int) int {
	pivot := (*sliceptr)[0]
	pivotIndex := 0
	n := len(*sliceptr)
	var swap int

	for i := n - 1; i >= 0; i-- {
		//number is greater than pivot and it is to the left of pivot, then swap
		if (*sliceptr)[i] > pivot && i < pivotIndex {
			swap = (*sliceptr)[i]
			(*sliceptr)[i] = pivot
			(*sliceptr)[pivotIndex] = swap
			pivotIndex = i
		}

		//number is smaller than pivot and it is to the right of pivot, then swap
		if (*sliceptr)[i] < pivot && i > pivotIndex {
			swap = (*sliceptr)[i]
			(*sliceptr)[i] = pivot
			(*sliceptr)[pivotIndex] = swap
			pivotIndex = i
		}
	}

	return pivotIndex
}
