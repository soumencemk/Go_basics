package main

import "fmt"

func main() {
	var i int
	for i = 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Print(i, " ==> EVEN\n")
		} else {
			fmt.Print(i, " ==> ODD\n")
		}
	}
}
