package main

import "fmt"

func main() {
	var x [5]int
	x[4] = 5
	fmt.Println(x)
	fmt.Println("the element at 5th index is:", x[4])
}
