package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		go f(0)
	}
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		return
	}
	var c chan string = make(chan string, 5) // buffer size of 5. Making the chan async

	go pinger(c)
	go ponger(c)
	go printer(c)

	_, err = fmt.Scanln(&input)
	if err != nil {
		return
	}
	// next
}

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
		amt := time.Duration(rand.Intn(255))
		time.Sleep(time.Millisecond * amt)
	}
}

func pinger(c chan<- string) {
	for i := 0; ; i++ {
		c <- "ping"
	}

}

func printer(c <-chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func ponger(c chan<- string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}
