package main

import (
	"fmt"
	"reflect"
)

func main14() {
	var typeDefinition interface{}
	typeDefinition = "sda"
	// Стандартный способ определения через switch
	switch i := typeDefinition.(type) {
	case string:
		fmt.Println("Это строка!", i)
	case int:
		fmt.Println("Это тип int!", i)
	default:
		fmt.Printf("Неизвестный тип: %T", typeDefinition)
	}
}

// Ну и с помощью reflect.TypeOf
func another() {
	var i interface{}
	i = 8
	t := reflect.TypeOf(i)
	fmt.Print(t)

}
