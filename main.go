package main

import "fmt"

// Type aliances
type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	userNames := make([]string, 2, 5)

	userNames[0] = "Daisy"

	userNames = append(userNames, "Max")
	userNames = append(userNames, "Emily")
	fmt.Println(userNames)

	couseRatings := make(floatMap, 3)
	couseRatings["go"] = 4.7
	couseRatings["flutter"] = 4.9
	couseRatings["django"] = 4.5
	fmt.Println(couseRatings)

	// with type aliaces - custom type, functions
	couseRatings.output()

	// For loops with arrays, slices, maps
	for index, value := range userNames {
		fmt.Println("Index:", index)
		fmt.Println("Value:", value)
	}

	for index, value := range couseRatings {
		fmt.Println("Key:", index)
		fmt.Println("Value:", value)
	}

}
