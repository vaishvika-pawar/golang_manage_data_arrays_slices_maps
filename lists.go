package main

import "fmt"

type Product struct {
	title string
	id    string
	price float64
}

func main() {
	var productNames [4]string = [4]string{"A Book"}
	prices := [4]float64{10.99, 9.99, 45.99, 70.0}
	fmt.Println(prices)

	productNames[2] = "A Carpet"
	fmt.Println(productNames)
	fmt.Println(prices[0])
	featuredPrices := prices[1:3]
	fmt.Println(featuredPrices)
	highlightedPrices := featuredPrices[:1]
	fmt.Println(len(highlightedPrices), cap(highlightedPrices))

	dynamicArray := []float64{10.99, 8.99}
	fmt.Println(dynamicArray)
	dynamicArray = append(dynamicArray, 25.00)
	fmt.Println(dynamicArray)
}
