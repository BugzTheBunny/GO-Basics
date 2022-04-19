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
