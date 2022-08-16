package main

import "fmt"

// Создаем родительскую структуру

type Human struct {
	age  uint8
	name string
}

func (h *Human) getAge() uint8 {
	return h.age
}

//Создаем структуру, с встраиваемой структурой Human

type Action struct {
	Human
	someTypeOfAction string
}

func main1() {
	//Наша структура унаследовала поля Human
	newAction := Action{Human{31, "James"}, "run"}
	//И мы можем обращаться к ним
	fmt.Println(newAction)
	fmt.Println(newAction.getAge())
	fmt.Println(newAction.name)
	fmt.Println(newAction.someTypeOfAction)
}
