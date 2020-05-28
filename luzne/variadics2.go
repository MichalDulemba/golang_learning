package main

import (
"fmt"
)

func main() {
testowaTablica := []int{2,3,4,2,344,5,5,3,3,4,4323,4,4}
x:=variadic(testowaTablica...)
fmt.Println(x)
}

func variadic(x...int)int{
for i, element:= range x{
fmt.Println(i, element)
}
return 0
}
