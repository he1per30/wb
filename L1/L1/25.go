package main

import (
	"fmt"
	"time"
)

func main25() {
	start := time.Now()
	ourSleep2(5)
	duration := time.Since(start)
	fmt.Println(duration)

}

func ourSleep2(second int) {
	<-time.After(time.Second * time.Duration(second))
}

// Самый не надежный вариант)
func ourSleep(second int) {
	for i := 0; i < second*2550000; i++ {

	}
}

// Добавляем в мапу каждый символ и если он повторяется return false
