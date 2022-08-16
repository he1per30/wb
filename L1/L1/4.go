package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main4() {
	var numberOfWorker int

	// Ввод n воркеров
	fmt.Print("Выберите количество воркеров:")
	_, err := fmt.Fscan(os.Stdin, &numberOfWorker)
	if err != nil {
		log.Fatal(err)
	}
	/* Создаем 3 канала
	канал m - для общения горутин
	канал q - для завершения работы всех воркеров
	канал done - ожидает завершение программы

	*/
	m := make(chan int, 10)
	done := make(chan os.Signal, 1)
	q := make(chan struct{})

	for i := 0; i < numberOfWorker; i++ {
		go worker(m)
	}

	go doWork(m, q)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-done
	//завершаем работу воркеров
	q <- struct{}{}
	close(m)

}

func doWork(m chan int, q chan struct{}) {
	i := 0
	for {
		select {
		case <-q:
			return
		default:
			m <- i
			i += 2
			time.Sleep(1 * time.Second)
		}

	}
}

func worker(jobs <-chan int) {
	for range jobs {
		fmt.Println(<-jobs)
	}
}
