package main

import "fmt"

func main() {
	var a, b int

	fmt.Scanf("%d", &a)
	fmt.Scanf("%d", &b)

	prod := a * b

	fmt.Printf("PROD = %v\n", prod)
}
