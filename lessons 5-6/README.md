# Arrays & Slices

- `Arrays` are declared like this:
```
...
func main() {
	var ages [3]int = [3]int{10, 20, 30}
	fmt.Println(ages, len(ages))
}
```
as expected, their length can't be changed once declared, they have one type, and you declare how many slots it has.
len(*) is a built in function that can return length of items, for example arrays or strings.

```
var ages [3]int = [3]int{10, 20, 30}
var name string = "Slava"
fmt.Println(ages, len(ages))
fmt.Println(name, len(name))
```

shorthand example:
```
func main() {
	names := [4]string{"Yoshi", "Mario", "Luigi", "Peach"}
	fmt.Println(names, len(names))
}
```
---
- ### `Slices` (uses array under the hood) are declared like this:
```
func main() {
	var scores = []int{100, 50, 200}
	fmt.Println(scores)
}
```
You do not need to specify how much size it will take.

Both `Arrays` and `Slices` are mutable, BUT, only `Slices` allows you to add items / remove items from it, by doing so:
```
func main() {
	var scores = []int{100, 50, 200}
	fmt.Println(scores)
	scores = append(scores, 300)
	fmt.Println(scores)
}
```
Note:
```
The resulting value of append is a slice containing all the elements of the original slice plus the provided values.
```
meaning that `append()` returns the list after its appended the new item.

---
- ### `Slice ranges`
Just like in other languages, you can select the range of items you want from a slice / array:
```
func main() {
	var scores = []int{100, 50, 200}
	scores = append(scores, 300)
	fmt.Println(scores[1:3]) # retuns [50 200]
	fmt.Println(scores[2:4]) # returns [200 300]
}
```
# The Standard Library

- `strings` library.

**`Go`** provides us with default libraries, for example `strings`, which gives us cool tools, for example the code below will check if the string (second argument is inside the first argument) and will return a boolean as a result, and some more cool things.
```
package main

import (
	"fmt"
	"strings"
)

func main() {
	greeting := "Hello World! I'm here!"
    // returns bool
	fmt.Println(strings.Contains(greeting, "Hello")) 

    // returns string with replaces items
	fmt.Println(strings.ReplaceAll(greeting, "Hello", "Hi")) 

    // returns string as upper case
	fmt.Println(strings.ToUpper(greeting)) 

    // returns index of string
	fmt.Println(strings.Index(greeting, "ll")) 

    // returns Slice by splitting by selected.
	fmt.Println(len(strings.Split(greeting, " ")), strings.Split(greeting, " ")) 
}
```

- `sort` library, the sort library allows us to interact with `Slices / Arrays` and gives cool features to use them.
```
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
	sort.Strings(names) // sorts
	fmt.Println(names) // prints sorted
	indexString := sort.SearchStrings(names, "Bowser") // finds index
	fmt.Println(indexString)
}

```