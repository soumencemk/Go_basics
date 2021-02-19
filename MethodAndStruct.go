package main

import "fmt"

//declared the structure named emp
type emp struct {
	name    string
	address string
	age     int
}

//Declaring a function with receiver of the type emp
func (e emp) display() {
	fmt.Println(e.name)
}

func main() {
	//declaring a variable of type emp
	var empdata1 emp

	//Assign values to members
	empdata1.name = "John"
	empdata1.address = "Street-1, Lodon"
	empdata1.age = 30

	//declaring a variable of type emp and assign values to members
	empdata2 := emp{
		"Raj", "Building-1, Paris", 25}

	//Invoking the method using the receiver of the type emp
	// syntax is variable.methodname()
	empdata1.display()
	empdata2.display()
}
