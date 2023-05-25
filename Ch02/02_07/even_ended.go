package main

import (
	"fmt"
)

// even ended number has same first and last digits.
// how many even ended numbers result from multiplying two 4-digit numbers?

func main() {
	// n := 42
	// // instead of printing to scrn, Sprintf returns a string
	// s := fmt.Sprintf("%d", n)

	// fmt.Printf("s = %q (type %T)\n", s, s)

	// var evenEndedNumCt int
	// iterate 1000 to 9999 with 2 for-loops
	// function takes in 2 numbers
	// multiply both
	// convert result to string
	// check if first char = last char
	// if yes, evenEndedNumCt++
	// finally, print evenEndedNumCt

	ct := 0

	for i := 1000; i < 10000; i++ {
		for j := 1000; j < 10000; j++ {
			p := j * i

			s := fmt.Sprintf("%d", p)
			if s[0] == s[6] {
				// fmt.Printf("%q is an even ended number\n", s)
				ct++
			}
		}
	}
	fmt.Printf("count: %d\n", ct)

}
