package main

import "fmt"

func main() {
	var a, b int

	fmt.Scanf("%d", &a)
	fmt.Scanf("%d", &b)

	fmt.Println(a) //printing just to debug
	fmt.Println(b) //printing just to debug

	produto := a * b

	fmt.Printf("PROD = %v\n", produto)
}
