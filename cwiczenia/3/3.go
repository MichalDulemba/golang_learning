package main

import (
	"fmt"
)

func main() {

	for i := 0; i <= 10000; i++ {
		fmt.Println(i)
	}

	z := 0
	for z < 10 {
		fmt.Println("z", z)
		z++
	}
	start := 1979
	for {

		fmt.Println(start)
		start++
		if start >= 2020 {
			break
		}

	}

	for i := 10; i <= 100; i++ {
		fmt.Println(i/4, i%4)

	}

	a := 3
	if a > 4 {
		fmt.Println("greater")
	} else {
		fmt.Println("smaller")
	}

	a = 44
	if a > 0 && a < 100 {
		fmt.Println("average weight")
	} else if a > 100 && a < 300 {
		fmt.Println("big weight")
	} else {
		fmt.Println("huge")
	}

	switch {
	case a > 0:
		{
			fmt.Println("larger than 33")
		}
	case a > 66:
		{
			fmt.Println("sadfsdfasadfsadf")
		}
	case a < 0:
		{
			fmt.Println("problem")
		}
	default:
		{
			fmt.Println("huge issue")
		}

	}

	favSport := "football"
	switch favSport {
	case "football":
		{
			fmt.Println("waste of time")
		}
	case "soccer":
		{
			fmt.Println("the same")
		}
	}

}
