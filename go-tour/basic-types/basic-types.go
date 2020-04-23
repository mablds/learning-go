package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe     bool       = false
	MaxInt   uint64     = 1<<64 - 1
	z        complex128 = cmplx.Sqrt(-5 + 12i)
	a        string     = "I think im in love with Golang"
	negative int        = -13
)

func main() {
	f := 3.142 //type infered to float
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
	fmt.Printf("Type: %T Value: %v\n", a, a)
	fmt.Printf("Type: %T Value: %v\n", negative, negative)
	fmt.Printf("Type: %T Value: %v\n", f, f)
}

//%T = type
//%v = value
