package main

import(
"fmt"
)

type returningIntFunction func() int

func main(){
 z := incr(10)
 fmt.Println(z())  
 fmt.Println(z())
 fmt.Println(z())
 fmt.Println(z())
 {
  // limiting scope
 c:=100
 fmt.Printf("%T \n",c)
  }
  
 //very similar to python generator
}

func incr(start int) returningIntFunction{
k:=start
return func()int{
k++
return k
}
}
