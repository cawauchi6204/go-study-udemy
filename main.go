package main

import "fmt"

func main() {
	a := 0
	if a == 2 {
		fmt.Println("two")
	} else {
		fmt.Println("I dont know")
	}

	if b := 100; b == 100 {
		fmt.Println("one hundred")
	}
}
