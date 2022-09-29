package main

import "fmt"

func main() {
	a := []int{1, 6, 3, 4, 5}
	fmt.Println(bigger(a))
}

func bigger(a []int) int {
	b := 0
	i := 0

	for i < len(a) {
		if b < a[i] {
			b = a[i]
		}
		i++
	}

	return b
}
