package main

import (
	"fmt"
	"time"
)

func routine(x string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(x)
	}
}
func main() {
	go routine("hello")
	routine("hi")
}
