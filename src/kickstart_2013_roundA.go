package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	var t int
	fmt.Scanf("%d", &t)

	for i := 1; i <= t; i++ {

		var n int
		fmt.Scanf("%d", &n)

		booksValueSlice := make([]int, n)

		for j := 1; j <= n; j++ {
			fmt.Scanf("%d", &booksValueSlice[j-1])
		}

		reslice := operation(booksValueSlice)
		fmt.Print("Case #", i, ": ")
		var restring []string
		for _,v:=range reslice{
			restring=append(restring, fmt.Sprint(v))
		}
		fmt.Print(strings.Join(restring," "))
		fmt.Println("")
		
	}
}

func operation(slint []int) []int {

	var oddslice []int
	var evenslice []int

	length := len(slint)

	for i := 0; i < length; i++ {
		if slint[i]%2 == 0 {
			evenslice = append(evenslice, slint[i])

		} else {
			oddslice = append(oddslice, slint[i])
		}
	}
	sort.Ints(oddslice)
	sort.Ints(evenslice)

	evenlength := len(evenslice)-1
	oddindex := 0

	for j := 0; j < length; j++ {

		if slint[j]%2 == 0 {
			slint[j] = evenslice[evenlength]
			evenlength--

		} else {
			slint[j] = oddslice[oddindex]
			oddindex++
		}
	}
	return slint
}