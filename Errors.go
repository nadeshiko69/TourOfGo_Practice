package main

import (
	"fmt"
)

type MyError struct {
	What float64
}

func (e *MyError) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e.What))
}

func Sqrt2(x float64) (float64, error) {
	if x < 0 {
		return 0, &MyError{x}
	}

	z := float64(1)
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z, nil
}

func errors() {
	fmt.Println(Sqrt2(2))
	fmt.Println(Sqrt2(-2))
}
