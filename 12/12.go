package main

import "fmt"

func main() {
	source := []string{"cat", "cat", "dog", "cat", "tree"}

	set := make(map[string]struct{})
	result := []string{}
	for _, value := range source {
		if _, ok := set[value]; !ok {
			set[value] = struct{}{}
			result = append(result, value)
		}
	}

	fmt.Println(result)
}
