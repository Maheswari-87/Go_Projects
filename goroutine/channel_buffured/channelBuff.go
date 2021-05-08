package main

import (
	"fmt"
)

/*func main() {
	c := make(chan int, 3)
	go func() {
		c <- 1 //Storing to channel
		c <- 2
		c <- 3
		c <- 4
		close(c)//After 4 values it stops iterating
	}()
	for i := range c {
		fmt.Println(i)
	}

}*/
type Car struct {
	Model string
}

func main() {
	c := make(chan *Car, 3)
	go func() {
		c <- &Car{"1"} //Storing to channel
		c <- &Car{"2"}
		c <- &Car{"3"}
		c <- &Car{"4"}
		close(c) //After 4 values it stops iterating
	}()
	for i := range c {
		fmt.Println(i.Model)
	}

}
