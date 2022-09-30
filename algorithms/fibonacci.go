package main

import "fmt"

func main() {
	Fibonacci(10)
}

func Fibonacci(number int) int {
	if number <= 1 {
		return number
	}

	// preencher as variáveis com os valores iniciais
	var penultResult, lastResult, result int
	penultResult, lastResult = 0, 1

	//interação para preencher até o index passado por parametro
	for currentIndice := 2; currentIndice <= number; currentIndice++ {
		result = lastResult + penultResult
		penultResult = lastResult
		lastResult = result
		fmt.Println(result)
	}

	return result
}
