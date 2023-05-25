package main

import (
	"fmt"
)

func main() {
	stocks := map[string]float64{
		"AMZN": 2087.98,
		"GOOG": 2540.85,
		"MSFT": 287.70, // must have trailing comma on multiline
	}

	fmt.Printf("length: %d\n", len(stocks))

	fmt.Printf("MSFT: %v\n", stocks["MSFT"])
	// zero if value not found
	fmt.Printf("TSLA: %v\n", stocks["TSLA"])

	value, keyfound := stocks["TSLA"]
	if !keyfound {
		fmt.Println("TSLA not found")
	} else {
		fmt.Println(value)
	}
	fmt.Println("adding TSLA")
	stocks["TSLA"] = 822.12
	fmt.Println(stocks)

	fmt.Println("deleting AMZN")
	delete(stocks, "AMZN")
	fmt.Println(stocks)

	for key := range stocks {
		fmt.Println(key)
	}

	for key, value := range stocks {
		fmt.Printf("%s -> %.2f\n", key, value)
	}
}
