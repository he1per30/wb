package main

import "fmt"

func main26() {
	fmt.Println(checkUniqueSting("abCdefAaf"))
}

// Добавляем в мапу каждый символ и если он повторяется return false
func checkUniqueSting(s string) bool {
	m := make(map[string]int)
	for _, elem := range s {
		if m[string(elem)] > 0 {
			return false
		}
		m[string(elem)]++
	}
	return true
}
