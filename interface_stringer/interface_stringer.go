package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)\n", p.Name, p.Age)
}
func main() {
	a := Person{"mahi", 23}
	b := Person{"nani", 25}
	fmt.Println(a, b)
}
