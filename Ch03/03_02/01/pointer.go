package main

import "fmt"

func main() {
	var x int = 42
	var p *int // Declare a pointer variable p of type *int

	p = &x // Assign the memory address of x to p

	fmt.Println("Value of x:", x)
	fmt.Println("Value of p:", p)
	fmt.Println("Value pointed by p:", *p) // Dereference p to get the value stored at that memory address

	*p = 24 // Modify the value of x indirectly through p

	fmt.Println("Updated value of x:", x)
}
