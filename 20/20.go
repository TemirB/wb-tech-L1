package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "snow dog sun"

	words := strings.Fields(s)
	for l, r := 0, len(words)-1; l < r; l, r = l+1, r-1 {
		words[l], words[r] = words[r], words[l]
	}

	fmt.Println(strings.Join(words, " "))
}
