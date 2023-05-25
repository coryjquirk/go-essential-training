package main

import (
	"fmt"
)

func main() {
	worker()
}

/*
the idiom of acquiring a resource, checking for an error,
and then deferring the release of the resource
is very common in Go
*/
func worker() {
	r1, err := acquire("A")
	if err != nil {
		fmt.Println("ERROR:", err)
	}
	defer release(r1)

	fmt.Println("worker")

	r2, err := acquire("B")
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	// defers the execution of function until the surrounding function returns
	// arguments evaluated immediately,
	// but function call not executed until surrounding function returns
	defer release(r2)
}

func acquire(name string) (string, error) {
	return name, nil
}

func release(name string) {
	fmt.Printf("Cleaning up %s\n", name)
}
