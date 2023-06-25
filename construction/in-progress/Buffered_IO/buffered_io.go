package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("movements.log")
	if err != nil {
		log.Fatal(err)
	}

	incoming, err := os.Create("outgoing.log")
	if err != nil {
		log.Fatal(err)
	}
	outgoing, err := os.Create("incoming.log")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	balance := 0
	linecount := 0

	for scanner.Scan() {
		// transactionType, err := incomingOrOutgoing(scanner.Text())

		// Implement parsing logic and balance incrementingd
		linecount++
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	fmt.Printf("Count: %d\n", linecount)
}

func incomingOrOutgoing(l string) (string, error) {
	// parse line to find the OUT IN keywords
	// return keyword, error
	// find out best way to propagate error from scanner
	return "NotImplemented", nil
}
