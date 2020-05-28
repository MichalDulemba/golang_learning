package main

import (
"fmt")

func main() {
    test1()
    fmt.Println("done")
}

func test1() {
    fmt.Println("start test1")
    defer fmt.Println("going out")
    fmt.Println("after defer - last part")
}
