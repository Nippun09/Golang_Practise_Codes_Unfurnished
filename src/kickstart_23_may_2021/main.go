package main

import (
	"fmt"
)

func main() {
	var T int
	var G int
	var RES int
	fmt.Scanf("%d", &T)
	for i := 1; i <= T; i++ {
		fmt.Scanf("%d", &G)
		RES = calcNumOfKs(G)
		fmt.Printf("Case #%d: %d", i, RES)
		fmt.Println("")
	}
}

func calcNumOfKs(g int) int {
	n := 1
	res := 0
	num := 2*g + n - (n * n)
	den := 2 * n

	for num > 0 {

		if num%den == 0 {
			res++
			n++
			num = 2*g + n - (n * n)
			den = 2 * n
		} else {
			n++
			num = 2*g + n - (n * n)
			den = 2 * n
		}
	}
	return res

}
