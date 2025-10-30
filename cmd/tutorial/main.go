package main

import "fmt"

func main() {
	equasionCount := 0
	fmt.Printf("How many equasions there's gonna be?:\n")
	fmt.Scan(&equasionCount)

	equasions, signs := getEquasions(equasionCount)
	fmt.Println(equasions)

	prices, base := getPrices(signs)
	fmt.Println(prices, base)
}
