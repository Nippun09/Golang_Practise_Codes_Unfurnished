package main

import "fmt"

func main() {
	/*
	   Change Machine
	   You have to make such a change machine that only returns the change in the form of coins.

	   You are supplied with an infinite number of quarters (25 cents), (10 cents), (5 cents), and (1 cent).
	   The user will enter any amount.
	   For each amount, you have to return the minimum number of coins possible!

	   Eg: if i enter the amount as 17 - it should give 10,5,1,1
	*/

	input := 25

	Payout(input)
}

func Payout(input int) {

	fixslice := []int{10, 25, 5, 1}

	counter := 0

	for {

		if input == 0 {
			break
		}

		if input >= fixslice[counter] {
			fmt.Print(fixslice[counter])
			fmt.Print(",")

			input = input - fixslice[counter]

		} else {

			counter++
			if input > fixslice[counter] {
				fmt.Print(fixslice[counter])
				fmt.Print(",")

				input = input - fixslice[counter]
			}

		}

	}

}
