package main

import "fmt"

func main() {
	bli := 14
	fmt.Printf("Type: %T Value: %v\n", bla, bla)
	fmt.Printf("Type: %T Value: %v\n", ble, ble)
	fmt.Printf("Type: %T Value: %v\n", bli, bli)
	fmt.Printf("Type: %T Value: %v\n", Blo, Blo)
	fmt.Printf("Type: %T Value: %v\n", Blu, Blu)
}

var bla int = 10
var ble = 12

const Blo, Blu = 16, 18

//Constantes não podem ser declaradas usando a sintaxe :=.
