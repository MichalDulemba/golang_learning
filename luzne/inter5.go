package main

import (
  "fmt"
)
// to tak nie dziala, tylko dla interfejsow mozna robic taka konwersje
type SciemnianyInt int
type SciemnianyFloat float32

func main(){
  var a SciemnianyFloat = 33.4
  b, ok := a.(SciemnianyInt)
  fmt.Println(b)
}
