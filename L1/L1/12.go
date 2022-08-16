package main

import "fmt"

func main12() {
	fmt.Println(convertToMap([]string{"cat", "cat", "dog", "cat", "tree"}))

}

// Просто добавляем в мапу и считаем количество животных
func convertToMap(s []string) map[string]int {
	mapCounts := make(map[string]int)
	for _, animal := range s {
		mapCounts[animal]++
	}

	return mapCounts
}
