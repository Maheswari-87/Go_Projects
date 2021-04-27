package main

import "fmt"

func fibonacci() func() int {
	one, two := 0, 1
	return func() int {
		ret := one
		one, two = two, one+two
		return ret
	}
}
func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
