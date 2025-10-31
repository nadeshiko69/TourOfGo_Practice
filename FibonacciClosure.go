package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	tmp1 := 0
	tmp2 := 1
	ret := 0
	return func() int {
		ret += tmp1
		tmp1, tmp2 = tmp2, ret
		return ret
	}
}

func fibonacciClosure() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
