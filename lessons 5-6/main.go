package main

import (
	"fmt"
	"sort"
)

func main() {
	// ---> Ints <---
	ages := []int{45, 20, 30, 99, 2, 52}
	/*
		This actually does not return a value,
		it sorts the current variable that is sent to the function
	*/
	sort.Ints(ages)
	fmt.Println(ages)

	// Finds the index of a value inside the array.
	index := sort.SearchInts(ages, 30)
	fmt.Println(index)

	// If element does not exist, it returns an index which is len + 1
	index2 := sort.SearchInts(ages, 999)
	fmt.Println(index2)

	// ---> Strings <---
	// All of the above functions work pretty much the same here.
	names := []string{"Yoshi", "Mario", "Luigi", "Bowser"}
	sort.Strings(names)                                // sorts
	fmt.Println(names)                                 // prints sorted
	indexString := sort.SearchStrings(names, "Bowser") // finds index
	fmt.Println(indexString)
}
