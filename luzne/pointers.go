package main

import(
"fmt")

func main(){
 z:=234
 fmt.Println(z)
 x:= &z
 *x=3333
 fmt.Println("address", &z)
 fmt.Println("z",z)
 P := *&z
 P = 234234234
 fmt.Println(P)
 fmt.Println("z",z)
 //fmt.Println("value", x)
 f:=3
 fmt.Println(f)
 changeStuff(&f)
 fmt.Println(f)
}

func changeStuff(p *int){

    *p=2342352352342343}
