package main

import (
	"fmt"
)

const (
	aaa int = 33
	vvv     = "ddd"
)
const (
	year = iota + 1992
	year1
	year2
	year3
)

func main() {

	const (
		llll int = 345345
	)

	x := 34
	s := fmt.Sprintf("%b", x)
	fmt.Println(s)

	s = fmt.Sprintf("%x", x)
	fmt.Println(s)

	s = fmt.Sprintf("%d", x)
	fmt.Println(s)

	a := 45 > 4
	fmt.Println(a)

	b := 34 >= 34
	fmt.Println(b)

	c := 4 < 44
	fmt.Println(c)

	f := "" != ""
	fmt.Println(f)

	const pleple = "AAAA"
	const newConst string = "cccc"

	l := 23445
	fmt.Printf("%d\t%b\t\t%x\n", l, l, l)
	l = l >> 1
	fmt.Printf("%d\t%b\t\t%x\n", l, l, l)

	newvar :=
		`firstline
		 secondline
	dkdkdkd
	dkdkdkdkdkdkdkd`
	fmt.Printf("\n%v\n\n", newvar)

	fmt.Println(year, year2, year3)
}
