package main

import (
	"fmt"
	"math"
)

func main10() {

	m := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	unionDegrees(m)

}

// Создаем мапу и проверяем значение к какому ключу оно относится
func unionDegrees(m []float64) {
	result := make(map[float64][]float64)
	for _, elem := range m {
		// Проверяем больше 0 или нет, чтобы понять в какую сторону округлять до 10х
		if elem < 0 {
			key := math.Ceil(elem/10) * 10
			result[key] = append(result[key], elem)
		} else {
			key := math.Floor(elem/10) * 10
			result[key] = append(result[key], elem)
		}
	}

	fmt.Println(result)

}
