package main

import "fmt"

func main() {
	fmt.Println(split(17))
}

//---------- Clean returns
func split(sum int) /*parametros */ (x, y int) /*retorno */ {
	x = sum * 4 / 9
	y = sum - x
	return //vai retornar os dois valores como inteiros
}
