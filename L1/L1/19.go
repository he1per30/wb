package main

import "fmt"

func main19() {
	fmt.Println(reverse("главрыба"))
}

// добавляем буквы в обратном порядке
func reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}
