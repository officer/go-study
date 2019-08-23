package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

func Add(t *tree.Tree, value int) {
	if t.Value == value {
		return
	}
	if value > t.Value {
		if t.Left != nil {
			Add(t.Left, value)
		} else {
			t.Left = &tree.Tree{nil, value, nil}
		}
	} else {
		if t.Right != nil {
			Add(t.Right, value)
		} else {
			t.Right = &tree.Tree{nil, value, nil}
		}
	}
}

func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for {
		x, okx := <-ch1
		y, oky := <-ch2
		fmt.Printf("x=%d y=%d\n", x, y)
		if x != y || okx != oky {
			return false
		}
		if okx == false {
			break
		}
	}
	return true
}

func main() {
	t1, t2 := tree.New(2), tree.New(2)
	Same(t1, t2)

}
