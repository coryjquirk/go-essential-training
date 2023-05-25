package main

import (
	"fmt"
)

// calculate maximal value in a slice

func main() {
	nums := []int{16, 8, 42, 4, 23, 15}
	fmt.Println(nums)

	max := 0

	for i := range nums {
		if nums[i] > max {
			max = nums[i]
		}
	}

	fmt.Printf("final max: %d", max)

}
