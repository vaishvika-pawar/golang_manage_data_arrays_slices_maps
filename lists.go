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

	// Task 1: Create new array (!) that contains
	// your three hobbies and output that array in the CL
	hobbies := [3]string{"Writing", "Singing", "Reading"}
	fmt.Println((hobbies))

	// Task 2: Also output more data about that array:
	// - The first element (standalone)
	// - The second and third element combined as a new list
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:3])

	// Task 3: Create a slice based on the first element
	// that contains the first and second elements
	// Create that slice in two different ways (i.e. create two slices in the end)
	mainHobbies := hobbies[0:2]
	fmt.Println(mainHobbies)

	// Task 4: Re-slice the slice from (3) and change it to contain
	// the secons and last element of the original array.
	fmt.Println(cap(mainHobbies))
	mainHobbies = mainHobbies[1:3]
	fmt.Println(mainHobbies)

	// Task 5: Create a "dynamice array" that containes your course goals (at least 2 goals)
	courseGoals := []string{"Learn Go!", "Learn all the basics"}
	fmt.Println(courseGoals)

	// Task 6: Set teh second goal to a different one And then add a third goal to that existing dynamice array
	courseGoals[1] = "Learn all the details!"
	courseGoals = append(courseGoals, "Excel in Flutter")
	fmt.Println(courseGoals)

	// Task 7: Bonus: Create a "Product" struct with title, id, price and
	// create a dynamic list of products (at least 2 products).
	// Then add a third product to the existing list of products.
	type Product struct {
		id    string
		title string
		price float64
	}
	products := []Product{{"P01", "Golang course", 799}, {"P02", "Flutter course", 885}}
	fmt.Println(products)
	newProduct := Product{
		"P03", "Python course", 499,
	}
	products = append(products, newProduct)
	fmt.Println(products)
}
