package main

import "fmt"

func main() {
	sl := []int{1, 2}
	fmt.Println(sl)

	sl4 := make([]int, 5)
	// makeは初期値を入れてくれる
	fmt.Println(sl4)

	sl4[1] = 1000
	fmt.Println(sl4)

	sl5 := []int{1, 2, 3, 4, 5}
	fmt.Println(sl5)

	fmt.Println(sl5[2:4])
}
