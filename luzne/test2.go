package main

import (
"fmt"
"strconv"
//"time"
//"os"
)

func main()  {
fmt.Println("new test")
fmt.Println(gimmename("aaaa",44.23))
}

func gimmename(a string, b float32) string {
	s32 := strconv.FormatFloat(float64(b), 'E', -1, 32)
	fmt.Printf("%T, %v\n", s32, s32)
    name := a+strconv.FormatFloat(float64(b), 'g', -1, 32)
    return name
}