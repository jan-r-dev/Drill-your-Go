package main

import (
	"fmt"
	"time"
)

func main() {
	quit := make(chan bool)

	go run_app("app_1", quit)

	time.Sleep(time.Duration(3 * time.Second))
	quit <- true

	fmt.Printf("Quit signal acknowledged. Received: %t\n", <-quit)
}

func run_app(name string, quit chan bool) {
	counter := 0
	for {
		select {
		case <-quit:
			fmt.Printf("%q received quit signal.\n", name)
			cleanup()
			quit <- true
			return
		default:
			fmt.Printf("%q doing some work. Cycle number:\t%d\n", name, counter)
			counter++
			time.Sleep(time.Duration(500 * time.Millisecond))
		}
	}

}

func cleanup() {
	fmt.Println("Performing cleanup operations...")
	time.Sleep(time.Duration(1 * time.Second))
	fmt.Println("Cleanup completed.")
}
