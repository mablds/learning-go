package main

import "fmt"

func main() {
	var a, b int
	var x int

	fmt.Scanf("%d", &a) // %d = base 10 integer
	fmt.Scanf("%d", &b) // %d = base 10 integer

	x = a + b

	fmt.Printf("X = %d\n", x)
}