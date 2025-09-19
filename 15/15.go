package main

import (
	"fmt"
	"strings"
)

var justString string

func createHugeString(size uint64) string {
	sb := strings.Builder{}
	sb.Grow(int(size))

	for i := uint64(0); i < size; i++ {
		sb.WriteByte('a' + byte(i%26))
	}

	return sb.String()
}

func someFunc() {
	// Создание большой строки на стеке, и выделение только первых 100 символов => не хорошо
	// v := createHugeString(1 &lt;&lt; 10) 1 << 10?
	// justString = v[:100]
	justString = createHugeString(100)
}

func main() {
	someFunc()

	fmt.Println(justString)
}
