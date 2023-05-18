package main

import (
	"fmt"
)

func main() {
	//basic
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}
	fmt.Println("break with auto-increment:")
	for i := 0; i < 3; i++ {
		fmt.Println(i)
		if i > 0 {
			break
		}
	}
	fmt.Println("continue:")
	for i := 0; i < 3; i++ {
		if i < 1 {
			continue
		}
		fmt.Println(i)
	}
	fmt.Println("manual increment:")
	a := 0
	for a < 3 {
		fmt.Println(a)
		a++
	}
	fmt.Println("break with manual increment:")
	b := 0
	for {
		if b > 2 {
			break
		}
		fmt.Println(b)
		b++
	}
}
