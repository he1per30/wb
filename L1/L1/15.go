package main

import "fmt"

var justString string

func someFunc() {

	v := createHugeString(1 << 10)
	// Наша задача проверить длину строки, чтобы не было паники
	if len(v) > 20 {
		justString = v[:20]
		fmt.Println(justString)
	}
	fmt.Println(len(v))

}

func createHugeString(i int) string {
	return "Не очень длинная строка, но для проверки пойдет"
}

func main15() {
	someFunc()
}
