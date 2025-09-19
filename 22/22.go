package main

import (
	"errors"
	"fmt"
	"math/big"
)

// Большие числа и операции
// Разработать программу, которая
// - перемножает,
// - делит,
// - складывает,
// - вычитает
// две числовых переменных a, b, значения которых > 2^20 (больше 1 миллион).

// Комментарий:
// в Go тип int справится с такими числами,
// но обратите внимание на возможное переполнение для ещё больших значений.
// Для очень больших чисел можно использовать math/big.

var ErrInvalidFormat error = errors.New("invalid format")

func parseStringToInt(s string) (*big.Int, error) {
	z, ok := new(big.Int).SetString(s, 10)
	if !ok {
		fmt.Printf("invalid format: %s\n", s)
		return nil, ErrInvalidFormat
	}

	return z, nil
}

func main() {
	sourceA := "123123123123123123123123123123123123123123123123123123123123123123"
	sourceB := "123123123123123123123123123123123123123123123123123123123123123123"

	a, errA := parseStringToInt(sourceA)
	if errA != nil {
		fmt.Printf("[ERROR] %s", errA.Error())
		return
	}

	b, errB := parseStringToInt(sourceB)
	if errB != nil {
		fmt.Printf("[ERROR] %s", errB.Error())
	}

	sum := new(big.Int).Add(a, b)
	diff := new(big.Int).Sub(a, b)
	prod := new(big.Int).Mul(a, b)

	quo, rem := new(big.Int), new(big.Int)
	quo.QuoRem(a, b, rem)

	rat := new(big.Rat).SetFrac(a, b) // точная дробь
	fmt.Printf("a = %s\nb = %s\n\n", a, b)
	fmt.Printf("a + b = %s\n", sum)
	fmt.Printf("a - b = %s\n", diff)
	fmt.Printf("a * b = %s\n", prod)
	fmt.Printf("a / b (quotient) = %s, remainder = %s\n", quo, rem)
	fmt.Printf("a / b (exact)    = %s\nApproximately %s\n", rat.RatString(), rat.FloatString(30))

}
