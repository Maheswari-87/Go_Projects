package main

import "fmt"

func Calculate(s int) int {
	y := s + 2
	return y
}
func main() {
	x := Calculate(2)
	fmt.Println("Result is ", x)
}
