package main

import (
	"fmt"
	"sort"
)

type Company struct {
  Name    string
  Address string
  Size    int
}
type bySize []Company
func (myvar bySize) Len() int { return len(myvar) }
func (a bySize) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a bySize) Less(i, j int) bool {return a[i].Size < a[j].Size }

func main() {
	//sorts inplace
	// s:=[]int{3,4,45,6,7,223,6,456,4,234,234,6}
	// sort.Ints(s)
	// fmt.Println("sorted",s)
	// sort.Reverse(s)
	// fmt.Println("reversed",s)
	s := []int{5, 2, 6, 3, 1, 4} // unsorted
	b := sort.IntSlice(s)
	fmt.Println("b", b)
	sort.Ints(sort.IntSlice(s))
	fmt.Println("sort na slice", s)
	//sort.Reverse(sort.Ints(sort.IntSlice(s)))
	//fmt.Println("reversed",s)
	sort.Sort(sort.Reverse(sort.IntSlice(s)))
	fmt.Println(s)

	//how to join slices
	X := []int{3, 4, 2, 4, 45, 6, 6, 6, 4, 4, 6, 6}
	Y := []int{6, 6, 7, 7, 7, 8, 8, 8, 54, 4}
	Z := append(X, Y...)
	fmt.Println(Z)



	companies := []Company{{"coke", "London", 234}, {"ibm", "ny", 3456}, {"hp", "ny", 343356}, {"ge", "la", 56}}
	fmt.Println(companies)
  sort.Sort(bySize(companies))
  fmt.Println(companies)
  sort.Sort(sort.Reverse(bySize(companies)))
  fmt.Println(companies)
}
