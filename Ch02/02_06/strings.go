package main

import (
	"fmt"
)

func main() {
	book := "The Catcher in the Rye"
	fmt.Println(book)

	fmt.Println(len(book))
	// uint84 is a byte:
	fmt.Printf("book[0] = %v (type %T)\n", book[0], book[0])
	// slice
	fmt.Println(book[4:11])
	// slice (no end)
	fmt.Println(book[4:])
	// slice (no start)
	fmt.Println(book[:11])
	// use + to concatenate
	fmt.Println("T" + book[1:])
	// unicode
	fmt.Println("½ price oranges")
	// multi line
	poem := `
	Down this road
	On a Monday morning
	Came a-riding
	Three strangers
	`
	fmt.Println(poem)
	// raw strings

}
