package main

import "fmt"

/*
	A declaração defer adia a execução de uma função até o final do retorno da função.
	Os argumentos das chamadas adiadas são avaliados imediatamente, mas a função chamada
	não é executada até o retorno da função.
*/
func main() {
	defer fmt.Println("Santos") // 5
	defer fmt.Println("dos")    //4
	defer fmt.Println("Lemos")  //3
	defer fmt.Println("Braga")  //2
	defer fmt.Println("Arthur") //1

	fmt.Println("Marcelo")

	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i) //empilhando defers
	}
}
