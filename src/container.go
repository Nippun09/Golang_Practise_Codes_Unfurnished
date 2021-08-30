/*
Container With Most Water

Given n non-negative integers a1, a2, ..., an , where each represents a point at coordinate (i, ai).
n vertical lines are drawn such that the two endpoints of the line i is at (i, ai) and (i, 0).
Find two lines, which, together with the x-axis forms a container, such that the container contains the most water.

Notice that you may not slant the container.

Input: height = [1,8,6,2,5,4,8,3,7]

Output: 49

Explanation: The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the container can contain is 49.

Example 2:

Input: height = [1,1]

Output: 1

Example 3:

Input: height = [4,3,2,1,4]

Output: 16

Example 4:

Input: height = [1,2,1]

Output: 2

*/

package main

import (
	"fmt"
	"math"
)

func main() {
	//1,8,6,2,5,4,8,3,7
	//4, 3, 2, 1, 4
	input := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}

	n := len(input)

	var maxvolume int

	for i := 1; i <= n; i++ {
		for j := 0; j <= n-i-1; j++ {
			fmt.Println("j: ", j, "j+i: ", j+i)
			min := math.Min(float64(input[j]), float64(input[j+i]))
			fmt.Println("min is: ", min)
			volume := int(min) * i
			fmt.Println("volume is: ", volume)
			if maxvolume < volume {
				maxvolume = volume
				fmt.Println("max volume is: ", maxvolume)
			}

		}

	}
	fmt.Println("maxvolume: ", maxvolume)
}
