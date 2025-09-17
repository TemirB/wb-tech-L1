package main

import (
	"fmt"
	"strconv"
)

// BitOperation представляет тип операции с битом
type BitOperation int

const (
	ClearBit BitOperation = iota // Установить бит в 0
	SetBit                       // Установить бит в 1
)

// printValue выводит значение в десятичном и двоичном формате с меткой времени
func printValue(timeLabel string, val int64) {
	fmt.Printf("[%s] decimal: %d \t binary: %s\n",
		timeLabel, val, strconv.FormatInt(val, 2))
}

// ModifyBit изменяет значение указанного бита в числе
// pos - позиция бита (0-based, где 0 - младший бит)
// operation - операция (ClearBit или SetBit)
func ModifyBit(v int64, pos uint, operation BitOperation) int64 {
	if operation == ClearBit {
		return v &^ (1 << pos)
	}
	return v | (1 << pos)
}

func main() {
	var value int64
	var bitPosition uint
	var operationInput int

	fmt.Print("Enter value, bit position, and operation (0-clear, 1-set): ")
	_, err := fmt.Scan(&value, &bitPosition, &operationInput)
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		return
	}

	// Валидация ввода операции
	if operationInput < 0 || operationInput > 1 {
		fmt.Println("Error: operation must be 0 (clear) or 1 (set)")
		return
	}

	operation := BitOperation(operationInput)
	result := ModifyBit(value, bitPosition, operation)

	fmt.Printf("\nBit manipulation results:\n")
	fmt.Printf("Value: %d, Position: %d, Operation: %s\n",
		value, bitPosition, []string{"Clear", "Set"}[operationInput])

	printValue("BEFORE", value)
	printValue("AFTER", result)

	// Пример работы:
	// Ввод: 5 1 1 = установить первый бит в числе 5 в значение 1
	//
	// [BEFORE] decimal: 5      binary: 101
	// [AFTER]  decimal: 7      binary: 111
	//
	// 7 =       1   1   1
	// 5 =       1   0   1
	// Позиции:  2   1   0
}
