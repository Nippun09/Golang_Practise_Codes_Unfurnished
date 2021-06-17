// Find the length of maximum number of consecutive numbers jumbled up in an array.
// Examples:
// Input : arr[] = {1, 94, 93, 1000, 5, 92, 78};
// Output : 3
// The largest set of consecutive elements is
// 92, 93, 94
// Input  : arr[] = {1, 5, 92, 4, 78, 6, 7};
// Output : 4
// The largest set of consecutive elements is
// 4, 5, 6, 7

package main

import (
	"fmt"
	"sort"
)

func main() {

	input := []int{1, 94, 93, 1000, 5, 92, 78}

	sort.Ints(input)

	n := len(input)

	//var arrlenmax int

	arrlenmax := 1

	largestlen := 0

	startindex := 0

	//var endindex int

	sequence := make(map[int]int)

	for i := 0; i < n-1; i++ {
		if input[i]+1 == input[i+1] {
			arrlenmax++
			//endindex=i+1
			if largestlen < arrlenmax {
				largestlen = arrlenmax
			}

		} else {
			sequence[arrlenmax] = startindex
			startindex = i + 1
			//endindex=i+1
			arrlenmax = 1
		}
	}

	fmt.Println(largestlen)

}
