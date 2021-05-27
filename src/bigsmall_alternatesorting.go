package main

import (
	"fmt"
	"sort"
)

func main() {
	//sorting
	//first elem : biggest
	//second: smalles
	//third : second biggest
	//fourth: second smallest
	//

	input := []int{78, 49, 99, 24, 56, 54, 23, 34, 12, 38, -10}
	var finalresult []int
	finalresult = append(finalresult, Sorter(input)...)
	//finalresult:=make(Sorter(input))
	fmt.Println(finalresult)

}

func Sorter(unsorted []int) []int {
	sort.Ints(unsorted)

	n := len(unsorted)

	var customsorted []int
	var smcount int
	var bigcount int
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			customsorted = append(customsorted, unsorted[n-1-bigcount])
			bigcount++
		} else {
			customsorted = append(customsorted, unsorted[smcount])
			smcount++
		}
	}
	return customsorted
}
