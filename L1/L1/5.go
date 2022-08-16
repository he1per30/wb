package main

import (
	"fmt"
	"os"
	"time"
)

func main5() {
	var timeLife time.Duration
	fmt.Print("Введите количество секунд: ")
	fmt.Fscan(os.Stdin, &timeLife)
	/* в переменной timeLife будет хранится количество секунду по истечению
	которых с помощью канала to завершается работа программы
	*/
	to := time.After(timeLife * time.Second)
	m := make(chan int, 100)
	go writeMessage(m)

	for {
		select {
		case mainChan := <-m:
			fmt.Println(mainChan)
		case <-to:
			return

		}
	}

}

func writeMessage(c chan<- int) {
	i := 1
	for {
		c <- i
		i++
		time.Sleep(1 * time.Second)
	}
}
