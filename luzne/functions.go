package main

import (
"fmt"
)
func main(){

testowaZmienna := "aaaaa"
fmt.Println("before",testowaZmienna)
terefere(&testowaZmienna)
fmt.Println("after", testowaZmienna)
fmt.Println(another())
fmt.Println(yetanother())
}

func terefere (argument *string) {
fmt.Println("test")
*argument = *argument + "bbbb"
return 
}

func another() []float64{
//a:= []int{2342,234,234,234}
return []float64{23.44442,234,234,234}
}
func yetanother() (float64, string){
return 234234.4,  "ssdfsdfsdf"
}
