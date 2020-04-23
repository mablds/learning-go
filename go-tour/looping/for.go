package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		sum += i
	}

	sum2 := 1
	for sum2 < 1000 { //While é feito em for sem ( ) mas sempre com { }
		sum2 += sum2
	}
	fmt.Println(sum)
	fmt.Println(sum2)

	//for { } //executará pra sempre. Looping infinito.
}
