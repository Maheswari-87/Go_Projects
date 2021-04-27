package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["answer"] = 32
	fmt.Println("the value is: ", m["answer"])

	m["answer"] = 42
	fmt.Println("the value is: ", m["answer"])

	delete(m, "answer")
	fmt.Println("The value is: ", m["answer"])

	v, ok := m["answer"]
	fmt.Println("The value:", v, "Present?", ok)

}
