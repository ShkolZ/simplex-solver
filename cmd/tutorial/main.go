package main

import (
	"fmt"
	"strings"
)

func main() {

	valuesMaps := make([]map[string]int, 3)
	for i := range valuesMaps {
		valuesMaps[i] = make(map[string]int)
	}
	priceMap := make(map[string]int)
	var signs [3]string
	var temp int

	equasions := make([]string, 0)

	eqNumber := 0
	fmt.Println("How many equasions there's gonna be?: ")
	fmt.Scan(&eqNumber)

	fmt.Println("Put some koefs in the place of '[]'")
	for i := 0; i < eqNumber; i++ {
		fmt.Println("[]x1 + ?x2 + ?x3 _ ?")
		fmt.Scan(&temp)
		valuesMaps[i]["x1"] = temp
		fmt.Printf("%vx1 + []x2 + ?x3 _ ?\n", valuesMaps[i]["x1"])
		fmt.Scan(&temp)
		valuesMaps[i]["x2"] = temp
		fmt.Printf("%vx1 + %vx2 + []x3 _ ?\n", valuesMaps[i]["x1"], valuesMaps[i]["x2"])
		fmt.Scan(&temp)
		valuesMaps[i]["x3"] = temp
		fmt.Printf("%vx1 + %vx2 + %vx3 [] ?\n", valuesMaps[i]["x1"], valuesMaps[i]["x2"], valuesMaps[i]["x3"])
		fmt.Scan(&signs[i])
		fmt.Printf("%vx1 + %vx2 + %vx3 %v []\n", valuesMaps[i]["x1"], valuesMaps[i]["x2"], valuesMaps[i]["x3"], signs[i])
		fmt.Scan(&temp)
		valuesMaps[i]["b"] = temp
		equasions = append(equasions, fmt.Sprintf("%vx1 +%vx2 + %vx3 %v %v\n", valuesMaps[i]["x1"], valuesMaps[i]["x2"], valuesMaps[i]["x3"], signs[i], valuesMaps[i]["b"]))
		fmt.Print(equasions[i])
		fmt.Println("--------------------------")
	}
	fmt.Printf("%v\n", equasions[0])
	fmt.Printf("%v\n", equasions[1])
	fmt.Printf("%v\n", equasions[2])

	generateMax(equasions, priceMap)

	for i, valMap := range valuesMaps {
		for j := 0; j < 3; j++ {
			if i == j {
				valMap[fmt.Sprintf("d%v", j)] = 1
			} else {
				valMap[fmt.Sprintf("d%v", j)] = 0
			}

		}
	}

	printMatrix(valuesMaps, priceMap)

}

func printMatrix(valuesMaps []map[string]int, priceMap map[string]int) {
	fmt.Printf("%v %v %v %v %v %v", priceMap["x1"], priceMap["x2"], priceMap["x3"], priceMap["d1"], priceMap["d2"], priceMap["d3"])
	fmt.Printf("x1 x2 x3 d1 d2 d3 b Omega")
	fmt.Println()
}

func generateMax(equasions []string, priceMap map[string]int) {
	var values [3]int
	finalString := ""

	fmt.Println("Insert values")
	fmt.Println("?x1 + ?x2 + ?x3")
	fmt.Scan(values[0], values[1], values[2])

	priceMap["x1"] = values[0]
	priceMap["x2"] = values[1]
	priceMap["x3"] = values[2]

	finalString = fmt.Sprintf("%vx1 + %vx2 + %vx3", priceMap["x1"], priceMap["x2"], priceMap["x3"])

	for i, equasion := range equasions {
		if strings.Contains(equasion, "=>") {
			finalString += fmt.Sprintf(" + 0d%v", i)
		}

	}

	fmt.Println(finalString)

}
