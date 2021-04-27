package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(
		strings.Contains("test", "es"),
		strings.Count("test", "t"),
		strings.HasPrefix("test", "te"),
		strings.HasSuffix("test", "st"),
		strings.ToLower("TEST"),
		strings.ToUpper("test"),
		strings.Index("test", "e"),
		strings.Join([]string{"a", "b"}, "-"),
		strings.Split("a-b-c-d-e", "-"),
		strings.Replace("aaaa", "a", "b", 2),
	)
}
