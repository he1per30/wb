package main

import (
	"fmt"
	"strings"
)

func main20() {
	fmt.Println(reverseWords("snow dog sun"))
}

// Как и в 19, только теперь добавляем слова в обратном порядке из массива
func reverseWords(s string) (result string) {
	words := strings.Fields(s)

	for _, elem := range words {
		result = elem + " " + result
	}
	return result
}
