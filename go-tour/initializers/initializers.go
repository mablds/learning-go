package main

import "fmt"

func main() {
	var i, j int = 1, 2                   //inicializadores em continuidade.
	fmt.Println(i, j, c1, python1, java1) //inicializadores sem tipar fortemente
}

var c1, python1, java1 = true, false, "no!" //inicializadores sem tipar fortemente - As atribuições já tipam a variável.
