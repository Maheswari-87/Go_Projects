package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "hello"
	a[1] = "world"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [10]int{0, 1, 2, 3, 5, 7, 9, 11, 13, 17}
	fmt.Println(primes)
}
