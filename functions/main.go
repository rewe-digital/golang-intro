package main

import "fmt"

func swap(x string, y string) (string, string) {
	return y, x
}

func main() {
	a := "AAA"
	b := "BBB"
	c, d := swap(a, b)
	fmt.Println(a, b)
	fmt.Println(c, d)
}
