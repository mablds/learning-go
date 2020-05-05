package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	c  = &Vertex{1, 2} // has type *Vertex
)

func main() {
	v := Vertex{1, 2}
	fmt.Println(v)

	v.X = 4
	fmt.Println(v)
	fmt.Println(v.X)

	p := &v
	p.X = 1e9
	fmt.Println(v)
	fmt.Println("--------------")
	fmt.Println(v1, c, v2, v3)
}
