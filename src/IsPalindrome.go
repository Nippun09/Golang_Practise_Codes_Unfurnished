package main

import (
	"fmt"
	"math"
)

func main() {

	input := 1234543210

	final := CheckPalindrome(input)

	fmt.Println("isPalindrome: ", final)

}

func CheckPalindrome(someint int) bool {
	var single, result int

	backup := someint

	for {

		single = someint % 10

		someint = someint / 10

		result = single + result*int(math.Pow10(1))

		if someint == 0 {
			break
		}
	}

	if backup == result {
		return true
	}

	return false
}
