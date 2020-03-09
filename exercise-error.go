package main

import (
	"fmt"
	"math"
)

type CustomError struct{
	ErrorType string
	Num float64
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("%v: cannot Sqrt negative number: %v", e.ErrorType, e.Num)
}

func Sqrt(x float64) error {
	z := 0.001
	if x < 0 {
		return &CustomError{
			ErrorType: "ValueError",
			Num: x,
		}
	} else {
		for math.Abs(z*z - x) > 0.0000001 {
			z -= 0.5 * ((z * z - x) / ( 2 * z ))
			fmt.Println(z)
		}
		return nil
	}
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
