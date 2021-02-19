package main

import "fmt"

func sample() {
	fmt.Println("Inside the sample()")
}

func printMe(a int) {
	fmt.Println(a)
}

func main() {
	//defer sample()
	defer printMe(1)
	defer printMe(2)
	defer printMe(3)
	defer printMe(4)
	fmt.Println("Inside the main()")
}
