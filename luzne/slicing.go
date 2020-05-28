package main

import (
"fmt"
"github.com/kr/pretty"
)
// install package with go get URL

func main(){
fmt.Println("nowy program")
dictionaryCalledMap := map[string]int{
"one":234234, "two":23, "four":234, "five":0,
}

a, ok := dictionaryCalledMap["seven"]
b, ok2 := dictionaryCalledMap["five"]
fmt.Println(a, ok,"\n\n")
fmt.Println(b, ok2)


//var anotherDictMap map[string]string
//anotherDictMap["eeee"]="ffff"
//fmt.Println(anotherDictMap)
fmt.Println(dictionaryCalledMap)
//dictionaryCalledMap["plepleple"]=234234


test := make(map[string]float64)
test["aaa"]= 234234.234234
test["sadkjf halskdj hlksadj fhlksad hfklasd fhklsajfhlksadfhlksajfhlksajd fhlksa dhfklsajd fhklsajd fhlksa jdhfklsajdhfklasjdfhlkasjd fh"] = 3.4

fmt.Println("testing loop")
var i int

for key, value := range test{
i++
fmt.Println(key, value)
}

for key, value:=range dictionaryCalledMap{
fmt.Println(key, value)
}


fmt.Println(test)

fmt.Println(len(test))

for i:=0; i<=calBle(10); i++{
fmt.Println(i)
}


newArray := []string{"asd","sfsdfg", "zxvczxcvxcv"}
for k, z:= range newArray{
fmt.Println(k, z)
}


type Z struct {
  width float64
  height float64
  desc map[string]string
}

var f Z
f.width =234
f.height = 7
f.desc = make(map[string]string)
f.desc["test"] =" passed"
f.desc["other attribute"]="not passed"
fmt.Println(f)

size := make(map[string] Z)
size["p"]= f

fmt.Printf("%# v",pretty.Formatter(size))



type operations struct {
   loops string
   ifs string
}
type functions struct {
   iterations string
   booleans string
   additional []string
}

type newlang struct {
    ops operations
    funcs functions
}

var bleble newlang
bleble.ops.loops = "test"
bleble.funcs.additional = append(bleble.funcs.additional, "sdfdfssdfa")

fmt.Printf("%# v", pretty.Formatter(bleble))

i=0
for i<=1000 {
fmt.Println("works forever",i)
i++
}
dziwnaTablica:=[]int{23,234,234}

for i, element :=range dziwnaTablica{
fmt.Println(i, element)
}

var nowaTablica[] int
nowaTablica = append(nowaTablica, 3939393939393)
for _, element :=range dziwnaTablica{
nowaTablica = append(nowaTablica, element)
}

fmt.Printf("%# v", pretty.Formatter(nowaTablica))
}
func calBle (myArg int)int{
   return myArg*myArg
}
