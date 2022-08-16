package main

import "fmt"

func main8() {

	n := changeBit(5, 0)
	fmt.Println(n)

}

// функция которая меняет бит на противоположный
func changeBit(n int64, pos uint) int64 {
	if hasBit(n, pos) {
		return clearBit(n, pos)
	}
	return setBit(n, pos)
}

// меняет заданный бит на 0
func clearBit(n int64, pos uint) int64 {
	mask := ^(1 << pos)
	n &= int64(mask)
	return n
}

// проверяет на 0 или 1
func hasBit(n int64, pos uint) bool {
	val := n & (1 << pos)
	return (val > 0)
}

// меняет заданный бит на 1
func setBit(n int64, pos uint) int64 {
	n |= (1 << pos)
	return n
}
