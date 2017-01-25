package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

// START OMIT
func (v Vertex) String() string {
	return fmt.Sprintf("[ x = %d | y = %d ]", v.X, v.Y)
}

func main() {
	v := Vertex{1, 2}
	fmt.Printf("v = %s", v)
}
// END OMIT
