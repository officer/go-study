package main

import (
	"fmt"
)

// Fundamental
type MyRoutine struct {
	x    int
	name string
}

func (m *MyRoutine) Add(x int) {
	for i := 0; i < 3; i++ {
		fmt.Printf("%s %d\n", m.name, m.x)
		m.x += x

	}
}

// Channel
func sum(s []int, c chan int) {
	sum := 0
	for _, num := range s {
		sum += num
	}
	c <- sum
}

// select
func selecters(c, end chan int) {
	for {
		select {
		case x := <-c:
			fmt.Printf("%d\n", x)
		case y := <-end:
			fmt.Printf("Quitting... %d\n", y)
		default:
			fmt.Println("Not yet")
		}
	}

}

func sendChannel(c chan int) {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for v := range nums {
		c <- v
	}
}

func selectTest() {
	x, y := make(chan int), make(chan int)
	go selecters(x, y)
	go sendChannel(x)
	sendChannel(y)
}

func main() {
	selectTest()
}
