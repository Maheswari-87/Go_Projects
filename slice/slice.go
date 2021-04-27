package main

import "fmt"

func main() {
	primes := [6]int{0, 1, 2, 3, 5, 7}

	var a []int = primes[1:4]
	fmt.Println(a)

}
