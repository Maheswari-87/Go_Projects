package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	x, y float64
}

func (v Vertex) sqrt() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}
func main() {
	a := Vertex{3, 4}
	fmt.Println(a.sqrt())
}
