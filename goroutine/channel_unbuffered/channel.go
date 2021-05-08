package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	go func() {
		c <- 1 //Storing to channel
	}()
	val := <-c
	fmt.Println(val)
	go func() {
		c <- 2 //Receiving from channel
	}()
	time.Sleep(time.Second * 2)
	val = <-c
	fmt.Println(val)
	fmt.Println(c) //printlng channel

}
