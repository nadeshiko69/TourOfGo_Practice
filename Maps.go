package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	for _, val := range strings.Fields(s) {
		m[val]++
	}

	// return map[string]int{"x": 1}
	return m
}

func maps() {
	wc.Test(WordCount)
}
