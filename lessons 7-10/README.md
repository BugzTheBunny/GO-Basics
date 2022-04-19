# Loops

Loops work just as in any other languages, here are examples of for loops.
Note that in the last example, im showing that you can't change the value of items inside the slice using the simple methods.

```
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
```

# Booleans & if else

`Booleans` and `if else` statments look normally:
```
func main() {
	age := 10
	fmt.Println(age > 10)  // age more then 10
	fmt.Println(age >= 10) // age more or equal to 10
	fmt.Println(age == 10) // age equal to 10
	fmt.Println(age < 10)  // age less then 10
	fmt.Println(age != 10) // age not equal to 10

	if age > 5 {
		fmt.Println("age is above 5")
	} else if age == 8 {
		fmt.Println("age is under 5")
	} else {
		fmt.Println("none of the above is true")
	}
}
```

# Functions
Note - you **`MUST`** specify what you return from the function, see in the examples how.
```
package main

import (
	"fmt"
	"strings"
)

// normal function.
func sayHello(name string) {
	fmt.Printf("Good morning %v \n", name)
}

// function with function argument.
func cycleNames(names []string, function func(string)) {
	for _, value := range names {
		function(value)
	}
}

// function that returns something.
func calculate(x int, y int) int {
	return x * y
}

// functions that returns multiple values.
func splitName(name string) (string, string) {
	splittedName := strings.Split(name, "-")
	return splittedName[0], splittedName[1]
}

func main() {
	sayHello("Slava")
	cycleNames([]string{"Slava", "Tom", "Josh"}, sayHello)
	result := calculate(123, 321)
	fmt.Println(result)
	first, last := splitName("Bugz-Bunny")
	fmt.Println(first, last)
}
```