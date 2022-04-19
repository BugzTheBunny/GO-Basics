# GO-Basics
Learning the basics of GO Language

_IMPORTANT BASIC Notes_:

1. Go is static language and is not interpreted language, meaning its super fast, and you must declare variable types.

2. to start the code you use 
    ```
    go run main.go
    ```



Examples:

----
Most basic app:
```
package main

import "fmt"

func main() {

	fmt.Println("Hello World!")
}
```

Declaring variables:
You can declare a type
```
var a int = 15
var x string = "This is text"
```
You can declare without explecitly saying that the variable is.
```
var y = "This is also text"
var z = 10
```

And you can use warlus operator.
```
new_text := "Rabbit"
```
Go has a wide variety of variable types:

----
## **DataTypes**

[Datatypes in GO](https://www.geeksforgeeks.org/data-types-in-go/)

```
ONLY POSITIVE:
uint8       the set of all unsigned  8-bit integers (0 to 255)
uint16      the set of all unsigned 16-bit integers (0 to 65535)
uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)
-----------
POSITIVE AND NEGATIVE:
int8        the set of all signed  8-bit integers (-128 to 127)
int16       the set of all signed 16-bit integers (-32768 to 32767)
int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)
-----------
float32     the set of all IEEE-754 32-bit floating-point numbers
float64     the set of all IEEE-754 64-bit floating-point numbers
-----------
complex64   the set of all complex numbers with float32 real and imaginary parts
complex128  the set of all complex numbers with float64 real and imaginary parts
-----------
byte        alias for uint8
rune        alias for int32
```

----
## **String formatting**

```
package main

import "fmt"

func main() {
	name := "Rabbit"
	age := 25
	fmt.Printf("My name is %v my age is %v \n", name, age)
}
```

creating a formatted string variable:
```
package main

import "fmt"

func main() {
	name := "Rabbit"
	age := 25
	formatted := fmt.Sprintf("My name is %v my age is %v \n", name, age)
	fmt.Println(formatted)
}
```