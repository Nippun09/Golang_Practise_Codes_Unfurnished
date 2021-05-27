package main

import "fmt"

func main() {
	intslice := []int{4, 5, 2, 99, 45, 23, 34, 12, 9, 89}

	finalresult := Sorter(intslice)

	fmt.Println("sorted array:", finalresult)
}

func Sorter(slint []int) []int {
	var swap int
	for i := 0; i < len(slint); i++ {
		for j := i; j < len(slint)-1; j++ {
			if slint[j] > slint[j+1] {
				swap = slint[j]
				slint[j] = slint[j+1]
				slint[j+1] = swap

			}
		}
	}

	return slint

}
