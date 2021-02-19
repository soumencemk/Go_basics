# GO LANG TUTORIAL

## First Program

```go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
}
```

* **package main** - Every Go Language program should start with a package name. Go allows us to use packages in another
  go programs and hence supports code reusability. Execution of a Go program begins with the code inside the package
  named main.

* **import fmt** - imports the package fmt. This package implements the I/O functions.

* **_func main()_** - This is the function from which program execution begins. The main function should always be
  placed in the main package. Under the main(), You can write the code inside { }.

* **fmt.Println** - This will print the text on the screen by the Println function of fmt.

## Data Types

* **Numeric types**
    * _int8_ - 8 bit signed integers.
    * _int16_ - 16 bit signed integers.
    * _int32_ - 32 bit signed integers.
    * _int64_ - 64 bit signed integers.
    * _uint8_ - 8 bit unsigned integers.
    * _uint16_ - 16 bit unsigned integers.
    * _uint32_ - 32 bit unsigned integers.
    * _uint64_ - 64 bit unsigned integers.
    * _float32_ - 32 bit floating point numbers.
    * _float64_ - 64 bit floating point numbers.
    * _complex64_ – has float32 real and imaginary parts.
    * _complex128_ - has float32 real and imaginary parts.
* **String Types**
    * Represents a sequence of bytes(characters). You can do various operations on strings like string concatenation,
      extracting substring, etc

* **Boolean types**
    * Two values - true/false

## Golang Interface

**Golang Interface** is a collection of method signatures used by a Type to implement the behavior of objects. The main
goal of Golang interface is to provide method signatures with names, arguments, and return types. It is up to a Type to
declare and implement the method. An interface in Golang can be declared using the keyword “interface.”

## Variables

**Variables** point to a memory location which stores some kind of value. The type parameter(in the below **_
Syntax :_** )
represents the type of value that can be stored in the memory location.

Variable can be declared using the **_Syntax :_**  :

```go
var x int8
```

> Once You declare a variable of a type You can assign the variable to any value of that type.

You can also give an initial value to a variable during the declaration itself using

```go
var x int = 5
```

If You declare the variable with an initial value, Go an infer the type of the variable from the type of value assigned.
So You can omit the type during the declaration using the **_Syntax :_**

```go 
var x = 10
```

Also, You can declare multiple variables with the **_Syntax :_**

```go
var i, j = 100, "Jake"
```

You can also declare a variable omiting the `var` keyword :

```go
a := 2
```

> Note that You used `:=` instead of `=`. You cannot use `:=` just to assign a value to a variable which is already declared. `:=` is used to declare and assign value.

> Variables declared without an initial value will have of 0 for numeric types, false for Boolean and empty string for strings

Refer to : [Variables.go](Variables.go) to see the examples in action.

## Constants

Constant in Go programming language is declared by using the keyword `const`

```go
const b = 10
fmt.Println(b)
b = 30 // Cannot reassign This will give error
fmt.Println(b) 
```

## Control structures and loops

Please see [Loops.go](Loops.go) for the implementations
> Go programming language supports only for loop

### For Loop

The **_Syntax :_**  of a for-loop :

```go
for initialisation_expression; evaluation_expression; iteration_expression{
// one or more statement
}

var i int
for i = 1; i <= 5; i++ {
fmt.Print(i, "\n")
}
```

### If-Else

The synax is

```go
if condition{
// statements_1
}else{
// statements_2
}
```

**for** and **if-else** together :

```go
for i = 1; i <= 10; i++ {
if (i%2==0){
fmt.Print(i, " ==> EVEN\n")
}else{
fmt.Print(i, " ==> ODD\n")
}
}
```

> if-else supports also chaining like java/c++

### Switch

**_Syntax :_**  :

```go
    a, b := 2, 1
switch a+b {
case 1:
fmt.Println("Sum is 1")
case 2:
fmt.Println("Sum is 2")
case 3:
fmt.Println("Sum is 3")
default:
fmt.Println("Printing default")
}
```

## Data structures

Please see [DataStructures.go](DataStructures.go) for the implementations

### Arrays

```go
var numbers [3]string //Declaring a string array of size 3 and adding elements 
numbers[0] = "One"
numbers[1] = "Two"
numbers[2] = "Three"
fmt.Println(numbers[1]) //prints Two
fmt.Println(len(numbers)) //prints 3
fmt.Println(numbers) // prints [One Two Three]

directions := [...]int{1, 2, 3, 4, 5} // creating an integer array and the size of the array is defined by the number of elements 
fmt.Println(directions)               //prints [1 2 3 4 5]
fmt.Println(len(directions)) //prints 5

//Executing the below commented statement prints invalid array index 5 (out of bounds for 5-element array)
//fmt.Println(directions[5]) 
```

#### Slice and Append Function

A **slice** is a portion or segment of an array. Or it is a view or partial view of an underlying array to which it
points. You can access the elements of a slice using the slice name and index number just as you do in an array. You
cannot change the length of an array, but you can change the size of a slice.

> Contents of a slice are actually the pointers to the elements of an array. It means **if you change any element in a slice, the underlying array contents also will be affected**.

**_Syntax :_**

```go
a := [5] string {"one", "two", "three", "four", "five"}
var b [] string = a[1:4] //created a slice named b
b[0] = "changed" // changed the slice data
```

_There are certain functions like Golang `len`, Golang `append` which you can apply on slices_

* `len(slice_name)` - returns the length of the slice
* `append(slice_name, value_1, value_2)` - Golang append is used to append value_1 and value_2 to an existing slice.
* `append(slice_nale1,slice_name2…)` – appends slice_name2 to slice_name1

```go

a := [5] string {"1", "2", "3", "4", "5"}
slice_a := a[1:3]
b := [5] string {"one", "two", "three", "four", "five"}
slice_b := b[1:3]

fmt.Println("Slice_a:", slice_a)
fmt.Println("Slice_b:", slice_b)
fmt.Println("Length of slice_a:", len(slice_a))
fmt.Println("Length of slice_b:", len(slice_b))

slice_a = append(slice_a, slice_b...) // appending slice
fmt.Println("New Slice_a after appending slice_b :", slice_a)

slice_a = append(slice_a, "text1") // appending value
fmt.Println("New Slice_a after appending text1 :", slice_a)
```

## Functions

**_Syntax_**

```go
func function_name(parameter_1 type, parameter_n type) return_type {
//statements
}
```

> The **parameters** and **return types** are optional. Also, you can return _multiple values_ from a function.

```go
package main

func main() {
	x, y := 15, 10
	sum, diff := sumAndDiff(x, y)
	print("SUM : ", sum, " DIFF : ", diff)
}

func sumAndDiff(x int, y int) (int, int) {
	return x + y, x - y
}
```

## Defer and stacking defers

`defer` statements are used to defer the execution of a function call until the function that contains the `defer`
statement completes execution.

```go
package main

import "fmt"

func sample() {
	fmt.Println("Inside the sample()")
}
func main() {
	//sample() will be invoked only after executing the statements of main()
	defer sample()
	fmt.Println("Inside the main()")
}
```

_Output :_

```bash
Inside the main()
Inside the sample()
```

> Here execution of `sample()` is deferred until the execution of the enclosing function(`main()`) completes.

Stacking defer is using multiple defer statements. Suppose you have multiple defer statements inside a function. Go
places all the deferred function calls in a stack, and once the enclosing function returns, the stacked functions are
executed in the Last In First Out(LIFO) order. You can see this in the below example.

```go
package main

import "fmt"

func display(a int) {
	fmt.Println(a)
}
func main() {
	defer display(1)
	defer display(2)
	defer display(3)
	fmt.Println(4)
}
```

Output :

```
4
3
2
1	
```

## Pointers

The `&` operator is used to get the address of a variable. It means `&a` will print the _memory address_ of variable a.

```go
a := 20
fmt.Println("Address:", &a)
fmt.Println("Value:", a)
```

A **pointer** variable stores the memory address of another variable. You can define a pointer using the syntax

```
var variable_name *type
```

```go
//Create an integer variable a with value 20
a := 20
	
//Create a pointer variable b and assigned the address of a
var b *int = &a

//print address of a(&a) and value of a  
fmt.Println("Address of a:", &a)
fmt.Println("Value of a:", a)

//print b which contains the memory address of a i.e. &a
fmt.Println("Address of pointer b:", b)

//*b prints the value in memory address which b contains i.e. the value of a
fmt.Println("Value of pointer b", *b)

//increment the value of variable a using the variable b
*b = *b+1

//prints the new value using a and *b
fmt.Println("Value of pointer b", *b)
fmt.Println("Value of a:", a)}
```
