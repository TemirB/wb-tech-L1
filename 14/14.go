package main

import (
	"fmt"
	"reflect"
)

func detectType(v interface{}) string {
	switch v.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	case chan interface{}:
		return "chan interface{}"
	case chan int:
		return "chan int"
	case chan string:
		return "chan string"
	case chan bool:
		return "chan bool"
	default:
		// Для более сложных случаев используем reflect
		t := reflect.TypeOf(v)
		if t != nil && t.Kind() == reflect.Chan {
			return fmt.Sprintf("chan %v", t.Elem())
		}
		return fmt.Sprintf("unknown type: %T", v)
	}
}

func main() {
	variables := []interface{}{
		42,
		"hello",
		true,
		make(chan int),
		make(chan string),
		make(chan bool),
		make(chan interface{}),
		3.14, // неизвестный тип
	}

	var myVar interface{}

	fmt.Scan(&myVar)
	variables = append(variables, myVar)

	for _, v := range variables {
		fmt.Printf("Значение: %v, Тип: %s\n", v, detectType(v))
	}
}
