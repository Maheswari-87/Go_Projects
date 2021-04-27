package main

import "fmt"

var mul = []int{0, 2, 4, 6, 8, 10, 12, 14, 16}

func main() {
	for i, v := range mul {
		fmt.Printf("2**%d=%d\n", i, v)
	}
}
