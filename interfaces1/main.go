package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X int
	Y int
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(float64(v.X * v.X + v.Y * v.Y))
}

func (v *Vertex) Move(dx int, dy int) {
	v.X += dx
	v.Y += dy
}

type Vertex3 struct {
	X int
	Y int
	Z int
}

func (v Vertex3) Abs() float64 {
	return math.Sqrt(float64(v.X * v.X + v.Y * v.Y + v.Z * v.Z))
}

// START OMIT
type Abser interface {
	Abs() float64
}

func compare(a Abser, b Abser) int {
	if a.Abs() == b.Abs() {
		return 0
	} else if a.Abs() > b.Abs() {
		return 1
	} else {
		return -1
	}
}

func main() {
	v := Vertex{1, 2}
	w := Vertex3{1, 2 ,3}
	fmt.Println(compare(v, w))
}
// END OMIT