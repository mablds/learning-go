package main

import "fmt"

func main() {
	Fibonacci(10)
}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	var penultResult, lastResult, result int
	penultResult, lastResult = 0, 1

	for curr := 2; curr <= n; curr++ {
		result = lastResult + penultResult
		penultResult = lastResult
		lastResult = result
		fmt.Println(result)
	}

	return result
}
