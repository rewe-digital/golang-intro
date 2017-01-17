package main

import (
	"fmt"
	"math"
)

// START OMIT
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

func main() {
	v := Vertex{1, 2}
	fmt.Println(v.Abs())

	v.Move(3, 4)
	fmt.Println(v)
}
// END OMIT