package main

import "fmt"

func main() {
	equasionCount := 0
	fmt.Printf("How many equasions there's gonna be?:\n")
	fmt.Scan(&equasionCount)

	equasions, signs := getEquasions(equasionCount)

	prices, base := getPrices(signs)
}
