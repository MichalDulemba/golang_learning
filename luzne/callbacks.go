package main

import (
"fmt"
"time"
)
type funkcjaObrazkowa func()
type XSTRUCT struct {
   a int
   b int
}
type funkcjaX func(int)string

func main(){
getData(callback)
v := XSTRUCT{3,4}
fmt.Printf("%T\n", v)
B(Z)
}


func Z(int)string{
   fmt.Println("aaaa")
   return "asdfasdf"
}

func B(P funkcjaX){
P(5)
}

func getData(A funkcjaObrazkowa){
time.Sleep(2 * time.Second)
A()
fmt.Printf("%T\n", A)
}


func callback(){
   fmt.Println("callback called") 
}
