package main

import (
	"fmt"
	"strings"
)

func getEquasions(eqNumber int, valuesMaps []map[string]int, signs [3]string) []string {
	temp := 0
	equasions := make([]string, 0)

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
	return equasions
}

func printMatrix(valuesMaps []map[string]int, priceMap map[string]int, baseValues *[3]string) {

	fmt.Printf("%-5v %-5v %-5v %-5v %-5v %-5v %-5v %-5v", "", "", priceMap["x1"], priceMap["x2"], priceMap["x3"], priceMap["d1"], priceMap["d2"], priceMap["d3"])
	for i, valueMap := range valuesMaps {
		fmt.Printf("%-5v %-5v %-5v %-5v %-5v %-5v %-5v %-5v", priceMap[baseValues[i]], baseValues[i], valueMap["x1"], valueMap["x2"], valueMap["x3"], valueMap["d1"], valueMap["d2"], valueMap["d3"])
	}

	fmt.Println()
}

func generateMax(equasions []string, priceMap map[string]int, baseValues *[3]string) {
	var values [3]int
	finalString := ""

	fmt.Println("Insert values")
	fmt.Println("?x1 + ?x2 + ?x3")
	fmt.Scan(&values[0], &values[1], &values[2])

	priceMap["x1"] = values[0]
	priceMap["x2"] = values[1]
	priceMap["x3"] = values[2]
	priceMap["d1"] = 0
	priceMap["d2"] = 0
	priceMap["d3"] = 0

	finalString = fmt.Sprintf("%vx1 + %vx2 + %vx3", priceMap["x1"], priceMap["x2"], priceMap["x3"])

	for i, equasion := range equasions {
		if strings.Contains(equasion, "=>") {
			finalString += fmt.Sprintf(" + 0d%v", i+1)
			baseValues[i] = fmt.Sprintf("d%v", i+1)
		}

	}

	fmt.Println(finalString)

}
