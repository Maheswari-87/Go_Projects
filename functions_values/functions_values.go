package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}
func main() {
	value := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(value(5, 12))
	fmt.Println(compute(value))
	fmt.Println(compute(math.Pow))
}
