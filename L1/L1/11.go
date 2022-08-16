package main

import (
	"fmt"
)

func main11() {
	numbers1 := []int{2, 4, 5, 11, 12, 8}
	numbers2 := []int{3, 1, 25, 5, 7, 8, 8}
	fmt.Println(intersect(numbers1, numbers2))
}

func intersect(nums1 []int, nums2 []int) []int {
	var intersectedNums []int
	numCounts := make(map[int]int)
	//Записываем в мапу значения, так чтобы у каждого ключа было значение равно 1
	for _, num := range nums1 {
		if _, ok := numCounts[num]; ok {
			numCounts[num]++
		} else {
			numCounts[num] = 1
		}
	}

	// Поиск по 2 массиву, если в маппе уже есть это значение, записываем его в массив
	// И для этого ключа делаем значение равно 0
	for _, num := range nums2 {
		if count := numCounts[num]; count > 0 {
			numCounts[num]--
			intersectedNums = append(intersectedNums, num)
		}
	}

	return intersectedNums
}
