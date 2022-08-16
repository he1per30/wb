package main

import "fmt"

func main6() {
	m := make(chan string)

	// передача канала в горутину и последующая запись в канал
	m <- "check"

	// Так же мы можем закрыть канал
	close(m)

	// с помощью return
	select {
	case <-m:
		return
	default:
		fmt.Println("Nothing available")
	}

}
