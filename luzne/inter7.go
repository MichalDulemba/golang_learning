package main

import (
  "fmt"
)

type ple struct{}

func main() {
  var i interface{} = ple{}
  switch i.(type){
  case int:
    fmt.Println("integer")
  case string:
    fmt.Println("string")
  default:
    fmt.Println("no idea")
  }

}
