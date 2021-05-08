package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(2) //added two functions
	func() {  //Anonymous function
		fmt.Println("Hello")
		wg.Done()
	}() //Calling the function
	func() { //Anonymous function
		fmt.Println("World")
		wg.Done()
	}()
	fmt.Println("Main")
	wg.Wait() //waiting for functions to execute
	fmt.Println("Exit")

}
