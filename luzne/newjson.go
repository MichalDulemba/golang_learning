// mozna tez uzywac Decoder / Encoder - podobno szybsze


package main

import (
"fmt"
"encoding/json"
"io/ioutil"
)

type person struct{
  First string `json:"first"`
  Last string `json:"Last"`
  Age int `json:"Age"`

}
func main (){
  p1 := person{
    First:"james",
    Last:"bond",
    Age: 23,
  }
  p2 := person{
    First: "Nadia",
    Last: "Dulemba",
    Age : 7,
  }
  people := []person{p1, p2}
  bs, err := json.Marshal(people)
  if err != nil{
    fmt.Println("blad", err)
  }
  fmt.Println("after conversion", string(bs))

  data, err := ioutil.ReadFile("./1.json")
  if err!=nil{
    fmt.Println("reading failed")
  }
  fmt.Println(data)
  var fixed person
  err = json.Unmarshal([]byte(data), &fixed)
  fmt.Println("error", err)
  fmt.Println(fixed)


  file, _ :=json.MarshalIndent(people, "", " ")
  err = ioutil.WriteFile("another.json", file, 0644)

  peopleFixed := []person{}
  backdata, err := ioutil.ReadFile("./another.json")
  err = json.Unmarshal([]byte(backdata), &peopleFixed)
  fmt.Println("people ", peopleFixed)
  fmt.Println("peopleFixed", peopleFixed[0].First)

  // saving and reading from map

  myDict := make(map[string]string)
  myDict["test"]="future"
  fmt.Println(myDict)
  data, err = json.Marshal(myDict)
  fmt.Println(data, err)

  newDict := make(map[string]string)
  err = json.Unmarshal(data, &newDict)
  fmt.Println("back", newDict)
  //var allDicts []myDict

}
