package main

import "fmt"

type total_app_state struct {
	killing       chan struct{}
	main_channel  chan string
	func_channel  chan func()
	current_state string
}
type easystruct struct {
	number int
}

var silnik *total_app_state

func main() {
	// a := easystruct{
	// 	number: 4,
	// }
	// fmt.Println(a)
	silnik = &total_app_state{
		killing:       make(chan struct{}),
		main_channel:  make(chan string),
		func_channel:  make(chan func()),
		current_state: "ted",
	}
	go superloop(silnik)
	testing(silnik)
	silnikInit(silnik)
	testing1000(silnik)
	fmt.Println(silnik)
}
func silnikInit(s *total_app_state) {
	s.current_state = "on"
}
func testing(state *total_app_state) {
	fmt.Println(state)
	state.main_channel <- "working"
	state.func_channel <- func() {
		fmt.Println("FROM FUNCTION testing BEFORE", state.current_state)
		state.current_state = "odododdodod"
		fmt.Println("FROM FUNCTION testing AFTER", state.current_state)
	}
	//state.killing <- struct{}{}
}
func testing1000(state *total_app_state) {
	fmt.Println(state)
	state.main_channel <- "offfff"
	state.func_channel <- func() {
		fmt.Println("testing1000 before", state.current_state)
		state.current_state = "fire!"
		fmt.Println("testing1000 after", state.current_state)
	}
	state.func_channel <- funcToSend
}

func funcToSend() {
	fmt.Println("func to send")
	fmt.Println(k)
	fmt.Println(state.current_state)
}
func superloop(state *total_app_state) {
	for {
		select {
		case z := <-state.main_channel:
			{
				fmt.Println(z)
			}
		case d := <-state.killing:
			{
				fmt.Println("killing it softly", d)
			}
		case g := <-state.func_channel:
			{
				fmt.Println("superloop", state.current_state)
				g()
				fmt.Println("superloop", state.current_state)
			}
		}

	}
}
