package main

import "fmt"

func main() {

	// AKA traditional "while" loop
	index := 0
	for index < 5 {
		fmt.Println("Value of index is: ", index)
		index++
	}

	// AKA traditional "for" loop
	for i := 0; i < 5; i++ {
		fmt.Println("Value of i is: ", i)
	}

	// AKA traditional "for" loop iterating over Slice
	names := []string{"Mario", "Luigi", "Bowser", "Yoshi"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	// AKA "for in" loop
	for ind, value := range names {
		fmt.Println(ind, value)
	}

	// AKA "for in" loop without using index
	for _, value := range names {
		fmt.Println(value)
	}

	// Example of how NOT to update values with iteration
	for _, value := range names {
		fmt.Println(value)
		value = "New Value" // This will not update the value inside 'names'
	}
	fmt.Println(names)

}
