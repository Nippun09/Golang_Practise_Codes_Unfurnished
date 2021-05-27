package main

import (
	"fmt"
	"math"
)

func main() {
	var T int
	var N, K int
	var S string
	var RES int
	fmt.Scanf("%d", &T)
	for i := 1; i <= T; i++ {
		fmt.Scanf("%d %d", &N, &K)
		fmt.Scanf("%s", &S)
		RES = calcNumOfPalins(S, N, K)
		fmt.Printf("Case #%d: %d", i, RES)
		fmt.Println("")
	}
}

func calcNumOfPalins(S string, N int, K int) int {
	res := 0
	upto := 0
	calc := 1
	index := 0

	if N%2 == 1 {
		upto = (N / 2) + 1
	} else {
		upto = N / 2
	}

	for i := 0; i < upto; i++ {
		if int(S[i]) == 97 {
			continue
		}
		calc = calc * (int(S[i]) - 97)
		index = i
		break
	}
	//
	rem := upto - 1 - index
	calc = calc * int(math.Pow(float64(K), float64(rem)))
	//
	if rem > 0 {
		calc = calc + (int(S[index+1]) - 97)
	}
	//
	if N > 4 {
		if int(S[upto-2]) < int(S[upto]) {
			calc = calc + 1
		}
	} else {
		if N > 1 {
			if int(S[upto-1]) < int(S[upto]) {

				calc = calc + 1
			}
		}

	}

	x := math.Pow(10, 9)
	y := int(x + 7)
	res = calc % y

	return res
}
