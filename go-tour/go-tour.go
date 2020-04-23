package main

import "fmt"

func main() {
	c2, python2, java2 := true, false, "noo!"
	fmt.Println(c2, python2, java2)
}

//---------- Variables

var c, python, java bool //Tipos no final. A declaração var pode estar num pacote ou a nível de função.
