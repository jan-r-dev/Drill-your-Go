package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := fanIn(report("Josh"), report("Jane"))

	counter := 0
	for report := range c {
		fmt.Println(report)
		counter++

		if counter == 10 {
			break
		}
	}

}

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)

	go func() {
		for {
			c <- <-input1
		}
	}()

	go func() {
		for {
			c <- <-input2
		}
	}()

	return c
}

func report(msg string) <-chan string {
	c := make(chan string)

	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s, %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()

	return c
}
