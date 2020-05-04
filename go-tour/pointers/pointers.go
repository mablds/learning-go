package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i         //point to i (42)
	fmt.Println(p)  //read i memory address through the pointer
	fmt.Println(*p) //read i through the pointer
	*p = 21         //set i through the pointer
	fmt.Println(&i) //see the memory address of i that is the same as p.
	fmt.Println(i)  //see the new value of i

	p = &j          //point to j
	fmt.Println(*p) //read j through the pointer
	fmt.Println(p)  //read p memory address that is the same as j through the pointer
	*p = *p / 37    //divide j through the pointer
	fmt.Println(j)  //see the new value of j
}
