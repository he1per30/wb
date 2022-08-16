package main

import (
	"fmt"
	"sync"
)

func main2() {

	wg := &sync.WaitGroup{}
	ourArray := []int{2, 4, 6, 8, 10}
	for i := range ourArray {
		wg.Add(1)
		go powerOfNumber(ourArray[i], wg)
	}
	wg.Wait()

}

// Простая функция с использованием WaitGroup
func powerOfNumber(someNumber int, wg *sync.WaitGroup) {
	fmt.Printf("квадрат числа: %d = %d\n", someNumber, someNumber*someNumber)
	wg.Done()
}
