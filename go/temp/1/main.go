package main

import (
	"fmt"
)

func pr(a ...int) {
	for i := range a {
		fmt.Print(a[i])
		if i != len(a) {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}

func main() {
	a := []int{1, 2, 3, 4, 5}
	pr(a...)
}
