package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		fmt.Printf("%d -> %f\n", i, z)
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func main() {
	actual := Sqrt(2)
	fmt.Printf("expect %f\n", math.Sqrt(2))
	fmt.Printf("actual %f\n", actual)
}
