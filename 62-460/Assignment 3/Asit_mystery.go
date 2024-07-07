package main

import "fmt"

// mystery is a function that takes a pointer to an integer as a parameter.
// It increments the value of the integer by 5 and returns the result of the expression 3*(value + 5) - 1.
func mystery(k *int) int {
	*k += 5
	return 3*(*k) - 1
}

// main is the entry point of the program.
// It calculates the sum1 and sum2 using the mystery function and prints the results.
// The values of i and j are initialized to 20.
// sum1 is calculated as (i / 2) + mystery(&i)
// sum2 is calculated as mystery(&j) + (j / 2)
// Finally, the values of sum1 and sum2 are printed to the console.
func main() {
	i, j := 20, 20
	sum1 := (i / 2) + mystery(&i)
	sum2 := mystery(&j) + (j / 2)

	fmt.Printf("sum1: %d\n", sum1)
	fmt.Printf("sum2: %d\n", sum2)
}
