package main

import (
"fmt"
)

func main(){
foo()
foo("aasdasd")
foo("ssdsdf", "zxcvzxcvxzcvxzcv")
}

func foo(tekst ...string){
   for i, element:= range tekst{
   fmt.Println(i, element)
}
}
