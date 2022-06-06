package main

import "fmt"

func uniqueNames(a, b []string) []string {
	temp := append(a, b...)
	check := make(map[string]int)
	var result []string = make([]string, 0)

	for _, val := range temp {
		check[val] = 1
	}

	for letter, _ := range check {
		result = append(result, letter)
	}
	return result
}

func main() {
	// should print Ava, Emma, Olivia, Sophia
	fmt.Println(uniqueNames(
		[]string{"Ava", "Emma", "Olivia"},
		[]string{"Olivia", "Sophia", "Emma"}))
}
