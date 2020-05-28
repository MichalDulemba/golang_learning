package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main(){
  fmt.Println("go")
  test(3245)
  // nie da sie nazwac argumentow przy wywolywaniu
  test23(5, "pleple234234")
  B := 21
  another1(&B)
  fmt.Println(B)
  start := time.Now()
  tabelka := createArray(10000000)
  end :=  time.Since(start)
  print("time1", end.Microseconds(), "\n")

  start1 := time.Now()
  test7(tabelka)
  end1 := time.Since(start1)
  print("time2", end1.Microseconds(), "\n")

  start2 := time.Now()
  test8(&tabelka)
  end2 := time.Since(start2)
  print("time3", end2.Microseconds(), "\n")
  test9("234234","sadfgsdfgsdfg","12safsdf")

}
func test9(testing ...string){
  for i, a:= range testing{
    fmt.Println(i, a)
  }
}

func test(a int){
  fmt.Println("a", a)
}

func test23(a int, b string ){
  var i int
  i = 0
  for i < a {
    i++;
    fmt.Println(b)
  }
}

func another1(a *int){
  *a = 33333
}

func createArray(N int) []int{
  a := make([]int, N)
  for i:=0; i<N; i++{
    a[i] = rand.Intn(100)
  }
  return a
}
func test7(a []int){
  fmt.Println(len(a))
}
func test8(a *[]int){
  fmt.Println(len(*a))
}
