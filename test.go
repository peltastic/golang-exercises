package main

import "fmt"

func main() {
	var x float64
	x = 20.0
	fmt.Println(x)
	fmt.Printf("x is of type %T\n", x)
	var s float64 = 20.0

	y := 42
	fmt.Println(s)
	fmt.Println(y)
	fmt.Printf("s is of type %T\n", s)
	fmt.Printf("y is of type %T\n", y)
	var a, b, c = 3, 4, "foo"

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("a is of type %T\n", a)
	fmt.Printf("b is of type %T\n", b)
	fmt.Printf("c is of type %T\n", c)
}
