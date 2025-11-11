package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// type Tree struct {
//     Left  *Tree
//     Value int
//     Right *Tree
// }

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	walk(t, ch) // メインの処理は別関数に切り出し、再帰呼び出しの途中でcloseするとpanicする
	close(ch)
}

func walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	walk(t.Left, ch)
	ch <- t.Value
	walk(t.Right, ch)
	// close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for v := range ch1 {
		if v != <-ch2 {
			return false
		}
	}
	return true
}

func equivalentBinaryTrees() {
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
