package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	x, y float64
}

func (v Vertex) abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}
func (v *Vertex) scale(f float64) {
	v.x = v.x * f
	v.y = v.y * f
}
func main() {
	f := Vertex{3, 4}
	f.scale(10)
	fmt.Println(f.abs())
}
