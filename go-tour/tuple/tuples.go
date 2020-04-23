package main

import "fmt"

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}

// -- Ok this is new to me --
func swap(x, y string) (string, string) {
	return y, x
}
