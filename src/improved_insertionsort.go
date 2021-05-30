package main

import (
	"fmt"
)

func main() {
	Input := []int{6, 5, 4, 2, 1, 10, 3, 7, 8, 9, 44, -2, -1}
	Output := Insertionsort(&Input)
	fmt.Println("sorted array is :", Output)
}

//improved insertion sort
//insertion sort has time complexity of O(n^2) and space complexity of O(1)
//but we can bring down time complexity to O(nlogn) if we increase
//space complexity from O(1) to O(n)
//reason for doing this: extra space used by algorithm can be reclaimed as soon as our program exits
//we will be using extra space for small amount of time, but out output generation will be fast
//rather than delaying for O(n^2) time, make use of extra space for some time i.e. O(nlogn) time
//and then reclaim the space.
//we are making use of binary search to find correct position of new element in the sorted array
//binary search has time complexity of O(logn), in worst case we will use binary search for each element i.e. n
//so time complexity of this algorithm is O(nlogn) and space complexity O(n)
func Insertionsort(input *[]int) []int {
	n := len(*input)
	var output []int
	output = append(output, (*input)[0])

	recently_inserted := (*input)[0]

	for i := 1; i < n; i++ {

		if (*input)[i] < recently_inserted {

			position := FindPositionByBinarySearch(output, (*input)[i])

			if position == 0 {

				intermediate := []int{(*input)[i]} //append([]int(), output[0:position]...)

				output = append(intermediate, output[position:]...)
			} else {

				intermediate := []int{(*input)[i]}
				intermediate = append(intermediate, output[position:]...)

				output = append(output[0:position], intermediate...)
			}

		} else {
			if (*input)[i] >= recently_inserted {

				output = append(output, (*input)[i])

				recently_inserted = (*input)[i]
			}

		}
	}
	return output
}

func FindPositionByBinarySearch(slint []int, elem int) int {
	n := len(slint)
	beg := 0
	end := n - 1
	var mid int

	mid = (beg + end) / 2

	for beg != end {

		if elem <= slint[beg] {
			return beg
		}
		if elem >= slint[mid] {
			beg = mid + 1
		}
		if elem <= slint[mid] {
			end = mid - 1
		}

		mid = (beg + end) / 2
	}

	if elem > slint[beg] {
		return beg + 1
	}
	return beg
}
