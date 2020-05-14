package main

import "fmt"

func main() {
	n := 3.14159
	var a, r float64

	fmt.Scanf("%f", &r) //%f for floating numbers but not long ones
	a = n * (r * r)

	fmt.Printf("A=%.4f\n", a)
}
