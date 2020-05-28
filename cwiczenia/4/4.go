package main

import (
	"fmt"
	"strconv"
)

func main() {
	myArr := [5]int{3, 4, 5, 22, 45}
	fmt.Println(myArr)
	for _, value := range myArr {
		fmt.Println(value)
	}
	fmt.Printf("%T", myArr)

	///////////

	anotherArr := []int{2, 4, 5, 22, 33, 44, 55, 66, 77, 88}
	fmt.Println(anotherArr)
	for i, value := range anotherArr {
		fmt.Println(i, value)
		fmt.Printf("%T \n", value)
	}
	fmt.Printf("%T \n\n", anotherArr)

	// a := []int{42,43,44,45,46}
	// b := []int{47,48,49,50,51}
	// c := []int{52,53,54,55,56}
	// d := []
	var slice1 []int
	var slice2 []int
	var slice3 []int
	var slice4 []int

	for i := 0; i < 5; i++ {
		slice1 = append(slice1, 42+i)
		slice2 = append(slice2, 47+i)
		slice3 = append(slice3, 44+i)
		slice4 = append(slice4, 43+i)
	}
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println(slice4)

	superArr := [11]int{34, 4, 4, 2, 2, 3, 4, 5, 2, 3, 4}

	sli1 := superArr[0:4]
	fmt.Println(sli1)

	sli2 := superArr[2:4]
	fmt.Println(sli2)

	sli3 := superArr[:]
	fmt.Println(sli3)

	sli4 := superArr[2:3]
	fmt.Println(sli4)

	sli5 := superArr[4:]
	fmt.Println(sli5)

	//////////////
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x, 52)
	fmt.Println(x)

	x = append(x, 53, 54, 55)
	fmt.Println(x)

	y := []int{56, 57, 58, 59}
	fmt.Println(y)

	x = append(x, y...)
	fmt.Println(x)

	// three dots in 4 places
	A := Sum(y...)
	fmt.Println(A)

	Total(3, 4, 4, 5)
	Total(3, 4)
	Total(y...)
	//Total(y) wrong

	textArr := []string{"aa", "bbb", "vvv"}
	fmt.Printf("%T \n", textArr)
	textArr2 := [...]string{"aaa", "ggg", "sadfasfdsafd", "vxvczzxvc"}
	fmt.Printf("%T \n\n", textArr2)

	//////////////////////////
	var k []int
	var k2 []int

	for i := 0; i <= 10; i++ {
		k = append(k, i+42)
	}
	fmt.Println(k)
	k2 = append(k[:3], k[5:]...)
	fmt.Println(k2)

	//////////////
	states := []string{` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`}
	fmt.Println(states)
	fmt.Println(len(states))
	fmt.Println(cap(states))
	for i := 0; i < len(states); i++ {
		fmt.Println(i, states[i])
	}

	l := make([]string, 0, 5)
	fmt.Println(l)
	fmt.Println(len(l))
	fmt.Println(cap(l))
	for i := 0; i < 128; i++ {
		l = append(l, "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"+strconv.Itoa(i))
	}
	fmt.Println(l)
	fmt.Println(len(l))
	fmt.Println(cap(l))

	/////////////////////////////////////////////////////
	o := [][]string{{"a", "b", "c"}, {"g", "v", "b"}}
	fmt.Println(o)
	fmt.Println(len(o))

	for _, u := range o {
		fmt.Printf("%T \n", u)
		for _, r := range u {
			fmt.Println(r)
		}
	}
	fmt.Println("\n\n\n\n\n")

	/////////////////////////

	var kosmos = make(map[string][]string)
	//kkk := []string{"udalo sie", "lub nie"}
	kosmos["test"] = []string{"udalo sie", "lub nie"}
	kosmos["test2"] = []string{"rakieta wyladowala", "albo sie rozbila", "albo jeszcze leci"}
	fmt.Println(kosmos)
	for k, v := range kosmos {
		fmt.Println(k, ":::::::")
		for index, element := range v {
			fmt.Println(index, element)
		}
	}

	space := map[string][]string{
		"ddd": []string{"ddd", "vvvv", "sfdsadsdf"},
	}
	fmt.Println(space)

	space["kuku"] = []string{"sdsdff", "sdxvvdfsa", "xzcvzxcv"}
	for key, value := range space {
		fmt.Println(key, value)
	}
	delete(space, "kuku")
	fmt.Println("\n\n", space)
}

func Sum(a ...int) int {
	totalSum := 0
	for _, element := range a {
		totalSum += element
	}
	return totalSum
}

func Total(a ...int) {
	fmt.Println(a)
}
