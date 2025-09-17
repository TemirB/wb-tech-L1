package main

import "fmt"

// Дана структура Human (с произвольным набором полей и методов).
// Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).

type Human struct {
	FirstName  string
	SecondName string
	isReady    bool
}

type Action struct {
	Human
}

func (h *Human) ReadyCheck() bool {
	return h.isReady
}

func main() {
	a := Action{
		Human{
			FirstName:  "Ivan",
			SecondName: "Ivanov",
			isReady:    true,
		},
	}

	if a.ReadyCheck() {
		fmt.Printf("Hello, %s\n", a.FirstName)
	} else {
		fmt.Println("Sorry, but you arent ready")
	}
}
