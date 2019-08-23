package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
	Scale(x, y float64)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	if v == nil {
		fmt.Println("NIL")
	}
	return math.Sqrt(math.Pow(v.X, 2) + math.Pow(v.Y, 2))
}

func (v *Vertex) Scale(x, y float64) {
	v.X *= x
	v.Y *= y
}

type IPAddr struct {
	a, b, c, d int
}

func (i IPAddr) String() (string, error) {
	if i.a < 0 || i.a > 255 {
		return fmt.Sprintf("Can not exceed 255 or below 0"), i
	}
	return fmt.Sprintf("%d.%d.%d.%d", i.a, i.b, i.c, i.d), nil
}

func (i IPAddr) Error() string {
	return fmt.Sprintf("Error")
}

type ErrorNavigation float64

func (e ErrorNavigation) Error() string {
	return fmt.Sprintf("%f", e)
}

func (e ErrorNavigation) Sqrt() (float64, error) {
	if 0 == 0 {
		return 0.1, ErrorNavigation(0.1)
	}
	return 0.1, nil
}

func main() {
	i := ErrorNavigation(5)
	fmt.Println(i.Sqrt())

}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
