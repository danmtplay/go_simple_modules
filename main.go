package main

import (
	"fmt"
)

func main() {
	// should print Ava, Emma, Olivia, Sophia
	names := merge.uniqueNames(
		[]string{"Ava", "Emma", "Olivia"},
		[]string{"Olivia", "Sophia", "Emma"})
	fmt.Println(names)

	// should print -1, -4
	x1, x2 := quadratic.findRoots(2, 10, 8)
	fmt.Printf("Roots: %f, %f", x1, x2)

}
