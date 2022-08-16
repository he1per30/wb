package main

import "fmt"

func main13() {
	a := 5
	b := 10
	// думаю тут комментарии не особо нужны)))
	a, b = b, a
	fmt.Printf("a = %d, b = %d\n", a, b)
}
