package main

import "fmt"

type MyError struct {
	number float64
}

func (e MyError) Error() string {
	return fmt.Sprintf("Invalid input %f", e)
}

type Vertex struct {
	x, y float64
}

func (v *Vertex) String() string {
	return fmt.Sprintf("%f %f\n", v.x, v.y)
}

func (v *Vertex) Scale(scale float64) error {
	if scale < 0 {
		return fmt.Errorf("Failed to scale %s")
	}
	v.x *= scale
	v.y *= scale
	return nil
}

func main() {
	v := Vertex{1.5, 1.5}
	fmt.Println(v)
	v.Scale(1.5)
	fmt.Println(v)
	e := v.Scale(-1)
	if e != nil {
		fmt.Printf("%v", e)
	}

}
