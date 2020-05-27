package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {

	fmt.Println("test")
	myChannel := make(chan string)
	go receiver(myChannel)

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go sender(myChannel)
		fmt.Println("launched", i)
	}
	wg.Wait()
}

func receiver(ch <-chan string) {
	i := 0
	counter := 0
	start := time.Now()
	for i < 2000 {
		value, state := <-ch
		if state {
			counter++
			fmt.Println("counter", counter, "value", value[:100], "state", state)
		}
		i++
	}
	end := time.Since(start)
	fmt.Println("this took", end.Seconds)
	wg.Done()
}

func sender(ch chan<- string) {
	i := 0
	for i < 1 {
		req, err := http.Get("http://www.google.com")
		fmt.Println(req.StatusCode)
		if err == nil && req.StatusCode == 200 {
			html, err := ioutil.ReadAll(req.Body)
			//https://haisum.github.io/2017/09/11/golang-ioutil-readall/
			//io.Copy(req.Body, html)
			if err != nil {
				fmt.Println("problem in IOUTIL", err)
				fmt.Println(req.Body)
				break
			} else if req.StatusCode == 200 {
				ch <- string(html)[:400]
				fmt.Println("html sent")
				i++
			}

		} else {
			fmt.Println("problem with request", err)
			return
		}

	}
	//close(ch)
	wg.Done()
}
