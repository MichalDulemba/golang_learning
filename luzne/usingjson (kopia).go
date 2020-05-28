package main
import (
"fmt"
"encoding/json"
"os"
)
type personInfo struct {
    FirstName string
    SecondName string
    Id int
    Age int
}

func main(){
newperson:=personInfo{FirstName:"John", SecondName:"Snow", Id:3, Age:23}
fmt.Println(newperson)
b, err := json.Marshal(newperson)
if err!=nil{
  fmt.Println("err", err)
}
fmt.Println("b", string(b))
os.Stdout.Write(b)
var fff personInfo
err=json.Unmarshal(b, &fff)
if err!=nil{
  fmt.Println("err", err)
}


 slcD := []string{"apple", "peach", "pear"}
    slcB, _ := json.Marshal(slcD)
    fmt.Println(slcB,string(slcB))



mapD := map[string]int{"apple": 5, "lettuce": 7}
    mapB, _ := json.Marshal(mapD)
    fmt.Println(string(mapB))

var c map[string]int
err=json.Unmarshal(mapB, &c)
fmt.Println(c)
}
