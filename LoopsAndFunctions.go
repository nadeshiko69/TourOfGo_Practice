package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := float64(1)
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func loopAndFunctions() {
	tmp := float64(12345)
	fmt.Println(Sqrt(tmp))
	fmt.Println(math.Sqrt(tmp))
}
