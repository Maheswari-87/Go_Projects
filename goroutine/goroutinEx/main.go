package main

import (
	"fmt"
	"time"
)

func heavy() {
	for {
		time.Sleep(time.Second * 1)
		fmt.Println("Heavy")
	}
}
func SuperHeavy() {
	for {
		time.Sleep(time.Second * 2)
		fmt.Println(" Super Heavy")
	}
}
func main() {
	go heavy()
	go SuperHeavy()
	fmt.Println("Main")
	select {}
}
