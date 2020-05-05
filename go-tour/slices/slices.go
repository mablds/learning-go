package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[0:6]
	fmt.Println(s)

	fmt.Println("-------------- slices are like pointers --------------")
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX" //all points to the names array. Slices are like references.
	fmt.Println(a, b)
	fmt.Println(names)

	fmt.Println("-------------- slice literals --------------")
	//A slice literal is like an array literal without the length.
	//And this creates the same array as above, then builds a slice that references it
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	n := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(n)

	fmt.Println("-------------- slice bounds --------------")
	z := []int{2, 3, 5, 7, 11, 13}

	z = z[1:4] //[2, 3, 5, 7, 11, 13] 1-3 => [3 5 7]
	fmt.Println(z)

	z = z[:2] // [3 5 7] 0-1 => [3 5]
	fmt.Println(z)

	z = z[1:] //[3 5] 1-1 => [5]
	fmt.Println(z)

	fmt.Println("-------------- slice length and capacity --------------")

	f := []int{2, 3, 5, 7, 11, 13}
	printSlice(f)

	// Slice the slice to give it zero length.
	f = f[:0]
	printSlice(f)

	// Extend its length.
	f = f[:4]
	printSlice(f)

	// Drop its first two values.
	f = f[2:]
	printSlice(f)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
