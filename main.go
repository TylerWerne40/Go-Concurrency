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
}

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
		amt := time.Duration(rand.Intn(255))
		time.Sleep(time.Millisecond * amt)
	}
}
