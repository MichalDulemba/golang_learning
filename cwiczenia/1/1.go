package main

import (
	"fmt"
	"reflect"
)

func main() {

	// Ninja 1
	x := 42
	y := "james bond"
	z := true

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	fmt.Println(x, y, z)

	var X int
	var Y string
	var Z bool

	fmt.Println(X, Y, Z)

	X = 42
	Y = "James Bond"
	Z = true
	// %v \t %v \t %v \t
	s := fmt.Sprintf("%v\t%v\t%v", X, Y, Z)
	fmt.Println(s)

	type Intowski int
	var d Intowski
	fmt.Println("d", d)
	fmt.Printf("%T \n", d)
	var mytype string
	mytype = fmt.Sprintf("%T", d)
	fmt.Println(mytype)
	d = 54
	fmt.Println("d", d)

	fmt.Println(reflect.TypeOf(d))
	fmt.Println(reflect.TypeOf(d).String())

	p := 333
	var c Intowski
	c = Intowski(p)
	fmt.Println(c, reflect.TypeOf(c))
}
