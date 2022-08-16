package main

import "fmt"

func main17() {
	a := []int{1, 2, 3, 6, 10, 15, 21, 28, 36, 45, 55}

	fmt.Println(binarySearch(a, 27))

}

func binarySearch(ourArray []int, items int) int {
	// проверка на наличие элементов в слайсе
	if len(ourArray) < 1 {
		return 0
	}
	// ищем середину
	middleNumber := (len(ourArray) - 1) / 2
	//сравниваем число с числом посередине и отсекаем половину
	if ourArray[middleNumber] > items {
		return binarySearch(ourArray[:middleNumber], items)
	} else if ourArray[middleNumber] == items {
		return items
	} else {
		return binarySearch(ourArray[middleNumber+1:], items)
	}
	return 0
}
