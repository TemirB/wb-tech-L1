package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "mygreatstring"
	runedString := []rune(s)
	var sb strings.Builder
	sb.Grow(len(s))

	for i := len(runedString) - 1; i >= 0; i-- {
		sb.WriteRune(runedString[i])
	}

	fmt.Println(sb.String())

}
