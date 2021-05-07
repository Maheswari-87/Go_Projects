package greeting

import "fmt"

func HelloM(s string) string {
	if len(s) == 0 {
		return "Hello dude!"
	}
	return fmt.Sprintf("Hello %v", s)
}
