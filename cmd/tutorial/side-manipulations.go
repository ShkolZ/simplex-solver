package main

import "fmt"

func getEquasions(equasionCount int) ([][4]int, []string) {
	equasions := make([][4]int, 0)
	signs := make([]string, 3)
	var equasion [4]int
	fmt.Println("Insert koeficients (put space in between):")
	for i := 0; i < equasionCount; i++ {
		fmt.Printf("%vx1 + %vx2 + %vx3 %v %v\n", "[]", "[]", "[]", "[]", "[]")
		fmt.Scan(&equasion[0], &equasion[1], &equasion[2], &signs[i], &equasion[3])
		equasions = append(equasions, equasion)
		fmt.Printf("%vx1 + %vx2 + %vx3 %v %v\n", equasion[0], equasion[1], equasion[2], signs[i], equasion[3])
		fmt.Printf("------------------------\n")

	}
	return equasions, signs
}

func getPrices(signs []string) ([]int, []int) {
	var answers [3]int
	prices := make([]int, 0)
	base := make([]int, 0)
	fmt.Println("Price equasion:")
	fmt.Println("[]x1 + []x2 + []x3 -> MAX")
	fmt.Scan(&answers[0], &answers[1], &answers[2])
	prices = append(prices, answers[0])
	prices = append(prices, answers[1])
	prices = append(prices, answers[2])
	finalStr := fmt.Sprintf("%vx1 + %vx2 + %vx3", prices[0], prices[1], prices[2])
	for i, sign := range signs {
		if sign == "<=" {
			prices = append(prices, 0)
			base = append(base, 0)
			finalStr += fmt.Sprintf(" + 0d%v", i)
		}
	}
	fmt.Println(finalStr)
	return prices, base
}
