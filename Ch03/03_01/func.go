package main

import (
	"fmt"
)

func main() {
	val := add(1, 2)
	fmt.Println(val)

	div, mod := divmod(4, 2)
	fmt.Printf("div=%d, mod=%d", div, mod)
}

func add(a int, b int) int {
	return a + b
}

// must close the return types in () when there are more than one
func divmod(a int, b int) (int, int) {
	return a / b, a % b
}
