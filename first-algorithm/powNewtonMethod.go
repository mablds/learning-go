package main

import "fmt"

func Sqrt(x float64) float64 {
	z := x / 2
	var oldValue float64 = 0.0

	for z*z != x { //while z times z isn't x...
		z -= (z*z - x) / (2 * z) //Newton's method
		fmt.Printf("Type: %T, Value %v\n", z, z)
		if z != oldValue { //if oldValue isn't the same as this actual z
			oldValue = z
		} else {
			return z
		}
	}

	return z
}

func main() {
	fmt.Println(Sqrt(25))
}
