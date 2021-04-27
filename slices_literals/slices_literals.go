package main

import "fmt"

func main() {
	i := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(i)

	b := []bool{true, false, true, true, false, true}
	fmt.Println(b)

	s := []struct {
		i int
		b bool
	}{
		{1, true},
		{2, false},
		{3, true},
		{4, true},
	}
	fmt.Println(s)
}
