package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}
type ByName []Person

func (c ByName) Len() int {
	return len(c)
}
func (c ByName) Less(i, j int) bool {
	return c[i].Name < c[j].Name
}
func (c ByName) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
func main() {
	kids := []Person{
		{"Jill", 9},
		{"Jack", 10},
	}
	sort.Sort(ByName(kids))
	fmt.Println(kids)
}
