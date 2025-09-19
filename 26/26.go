package main

import (
	"fmt"
	"strings"
)

func IsUnique(s string) bool {
	if s == "" {
		return true
	}
	s = strings.ToLower(s)
	seen := make(map[rune]struct{})
	for _, r := range s {
		if _, ok := seen[r]; ok {
			return false
		}
		seen[r] = struct{}{}
	}
	return true
}

func main() {

	cases := []string{
		"abcdA",
		"adeedjedenideo",
		"",
		"abcd",
	}

	for i, c := range cases {
		fmt.Printf("[#%d] ", i)
		if IsUnique(c) {
			fmt.Println("OK: все символы уникальны")
		} else {
			fmt.Println("Найдены дубликаты")
		}
	}
}
