package main

import (
  "fmt"
  //"string"
)


type Writer interface{
  Write([]byte) (int, error)
  // write -> writer, read-> reader
}

type ConsoleWriter struct {}
func (cw ConsoleWriter) Write(data []byte)(int, error){
  n, err := fmt.Println(string(data))
  return n, err
}

// func  drukarka(w Writer, a string) (int, error){
//   result, err := w.Write(a)
//   print("from drukarka", result, err)
//   return result, err
// }


func main() {
  var w Writer = ConsoleWriter{} // tu mozna zamienic na inne typy writerow - np tcp writer (zachowanie polimorficzne)
  w.Write([]byte("ple ple ple"))
}
