package main

import "fmt"

func main23() {
	k := []int{1, 2, 3, 4, 5}
	k = delete(k, 0)
	fmt.Println(k)
}

func delete(k []int, i int) []int {
	k = append(k[:i], k[i+1:]...)
	return k
}
