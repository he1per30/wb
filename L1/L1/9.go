package main

import (
	"fmt"
	"time"
)

func main9() {
	values := []int{1, 2, 3, 4, 5}
	// создаем 2 канала
	firstChan := make(chan int)
	secondChan := make(chan int)
	// и 3 функции под них, которые общаются между собой по средствам этих 2х каналов
	go writeInChan(values, firstChan)
	go readFromChanel(firstChan, secondChan)
	go printValues(secondChan)
	time.Sleep(1 * time.Second)
}

func writeInChan(m []int, c chan int) {
	for _, elem := range m {
		c <- elem
	}
}

func readFromChanel(readChan <-chan int, writeChan chan<- int) {
	for elem := range readChan {
		writeChan <- elem * 2
	}
}

func printValues(resultChan <-chan int) {
	for elem := range resultChan {
		fmt.Println(elem)
	}
}
