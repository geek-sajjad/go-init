package main

import "fmt"

func main() {
	// var x int
	// x = 10

	// y := 20

	// z := x + y

	// const test = "testing"

	// fmt.Println(z)

	// fmt.Println(add(1, 2))

	// if x > 5 {

	// } else {

	// }
	var array [3]int
	array[0] = 2

	slice := []int{1, 2, 45, 3}

	for i := 0; i < 4; i++ {
		fmt.Println(slice[i])
	}

}

func add(a, b int) int {
	return a + b
}
