package main

import "fmt"

func main() {
	var i int
	fmt.Print("Enter a number:")
	fmt.Scanf("%v", &i)
	switch i {
	case 0:
		fmt.Println("zero")
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	case 4:
		fmt.Println("four")
	case 5:
		fmt.Println("five")
	default:
		fmt.Println("unknown number")
	}
	/*i := 10
	if i > 10 {
		fmt.Println("Big")
	} else {
		fmt.Println("Small")
	} Prints small*/

}
