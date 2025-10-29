package main

import (
	"fmt"
)

func main() {

	valuesMaps := make([]map[string]int, 3)
	for i := range valuesMaps {
		valuesMaps[i] = make(map[string]int)
	}
	priceMap := make(map[string]int)
	var signs [3]string
	var baseValues [3]string
	var omega [3]int
	zjcj := make([]int, 0)

	eqNumber := 0
	fmt.Println("How many equasions there's gonna be?: ")
	fmt.Scan(&eqNumber)

	equasions := getEquasions(eqNumber, valuesMaps, signs)

	fmt.Printf("%v\n", equasions[0])
	fmt.Printf("%v\n", equasions[1])
	fmt.Printf("%v\n", equasions[2])

	generateMax(equasions, priceMap, &baseValues)

	for i, valMap := range valuesMaps {
		for j := 0; j < 3; j++ {
			if i == j {
				valMap[fmt.Sprintf("d%v", j)] = 1
			} else {
				valMap[fmt.Sprintf("d%v", j)] = 0
			}

		}
	}

	calculateZjcj(&zjcj)
	calculateOmega(&omega)
	printMatrix(valuesMaps, priceMap, &baseValues)

}
