package main

import (
	"fmt"
)

func main() {
	loons := []string{"bugs", "daffy", "taz"}
	fmt.Printf("loons = %v (type %T)\n", loons, loons)
	fmt.Printf("length: %d\n", len(loons))
	fmt.Println("=========")
	fmt.Println(loons[1])
	fmt.Println("=========")
	fmt.Println(loons[1:])
	fmt.Println("=========")

	for i := 0; i < len(loons); i++ {
		fmt.Println(loons[i])
	}

	for i := range loons {
		fmt.Println(i)
	}
	// double value range
	for i, name := range loons {
		fmt.Printf("%s at %d\n", name, i)
	}

	for _, name := range loons {
		fmt.Println(name)
	}

	loons = append(loons, "elmer")
	fmt.Println(loons)

}
