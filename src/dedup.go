package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {

	var str string
	var rslice []string

	//de-duplicating this slice without using map or sorting the slice
	input := []int{5, 4, 1, 3, 3, 2, 1}

	deduped := []int{}

	for _, v := range input {
		rslice = append(rslice, string(strconv.Itoa(v)))
	}
	str = strings.Join(rslice, "")

	fmt.Println(str)

	n := len(rslice)

	fmt.Println(n)

	z := len(str)

	fmt.Println("z: ", z)

	for i := 0; i < n-2; i++ {
		c := string(str[i])
		fmt.Println("c: ", c)
		intermediate := strings.Join(rslice[0:i], "")
		fmt.Println("intermediate: ", intermediate)
		interim := strings.ReplaceAll(strings.Join(rslice[i+1:], ""), c, "!")
		fmt.Println("interim: ", interim)
		str = intermediate + c + interim
		fmt.Println("str: ", str)
		rslice = strings.Split(str, "")

	}

	//rslice = strings.Split(str, "!")
	r := []rune(str)

	for i, v := range r {

		if unicode.IsNumber(v) {
			f, _ := strconv.Atoi(rslice[i])
			deduped = append(deduped, f)
		}

	}

	fmt.Println(deduped)
}
