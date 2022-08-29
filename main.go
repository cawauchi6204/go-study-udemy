package main

import "fmt"

func main() {
	var x interface{}
	fmt.Println(x)
	x = 1
	fmt.Println(x)

	// interface型は全ての方を汎用的に表す
	// 初期値はnil

	x = [3]int{1, 2, 3}
	fmt.Println(x)
}
