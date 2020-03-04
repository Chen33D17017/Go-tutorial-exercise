package main

import (
	"fmt"
	"math"
)


func Sqrt(x float64) float64 {
	z := 0.001
	for math.Abs(z*z - x) > 0.0000001 {
		z -= 0.5 * ((z * z - x) / ( 2 * z ))
		fmt.Println(z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(25))
}
