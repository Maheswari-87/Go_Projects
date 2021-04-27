package main

import "fmt"

func pointer(xPtr *int) {
	*xPtr = 0
}
func main() {
	x := 5
	pointer(&x)
	fmt.Println(x)
}
