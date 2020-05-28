package main

import (
"fmt"
"golang.org/x/crypto/bcrypt"
)


func main(){
  s:= "plepelpel234"
  bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
  if err != nil{
    fmt.Println("err",err)
  }
  fmt.Println(s)
  fmt.Println(bs)
}
