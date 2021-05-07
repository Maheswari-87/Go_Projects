package main

import (
	"Testing_1/greeting"
	"fmt"
)

func main() {
	p := greeting.HelloM("")
	fmt.Println(p)
	s := greeting.HelloM("Mahi")
	fmt.Println(s)
}
