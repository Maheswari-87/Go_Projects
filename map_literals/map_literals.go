package main

import "fmt"

type Vertex struct {
	Lan, Long float64
}

var m = map[string]Vertex{
	"bell labs": {40.68433, -74.39967},
	"google":    {37.42202, -122.08408},
}

func main() {
	fmt.Println(m)
}
