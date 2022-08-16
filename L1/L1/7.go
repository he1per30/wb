package main

import (
	"fmt"
	"sync"
	"time"
)

type testMap struct {
	m map[string]int
	*sync.Mutex
}

func main7() {
	newMap := &testMap{
		m: make(map[string]int),
	}
	// Используем мьютексы
	for i := 0; i < 10; i++ {
		go newMap.writeInMap("test")
	}
	// Используем sync.Map
	var counters sync.Map
	for i := 0; i < 5; i++ {
		go writeInMapSyncMap(&counters, i)

	}
	time.Sleep(4 * time.Second)

}

// Запись с использованием мьютексов sync.Map

func writeInMapSyncMap(m *sync.Map, value int) {
	m.Store("test", value)
	fmt.Println(m.Load("test"))
}

// Запись с использованием мьютексов

func (ourMap *testMap) writeInMap(key string) {
	ourMap.Mutex.Lock()
	defer ourMap.Mutex.Unlock()
	ourMap.m[key]++
	//fmt.Printf("For key: %s, value = %d\n", key, ourMap.m[key])

}
