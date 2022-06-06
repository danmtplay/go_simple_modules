package merge

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
