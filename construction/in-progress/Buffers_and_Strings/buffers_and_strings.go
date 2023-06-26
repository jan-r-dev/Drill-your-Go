package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type Transaction struct {
	Name    string
	Type    string
	Amount  int
	isValid bool
}

func (t *Transaction) populateFromLog(movement string) {
	s := strings.Split(movement, ",")

	if len(s) < 3 {
		t.isValid = false
		return
	}

	amount, err := strconv.Atoi(s[2])
	if err != nil {
		t.isValid = false
		return
	}

	t.Name = s[0]
	t.Type = s[1]
	t.Amount = amount
	t.isValid = true

	t.isTransactionValid()
}

func (t *Transaction) isTransactionValid() {

	if t.Name == "" {
		t.isValid = false
	}

	if t.Type != "OUT" && t.Type != "IN" {
		t.isValid = false
	}

	if t.Amount < 0 {
		t.isValid = false
	}
}

func (t *Transaction) writeToLog(w io.Writer) error {
	movement := fmt.Sprintf("%s,%s,%d\n", t.Name, t.Type, t.Amount)

	_, err := w.Write([]byte(movement))

	return err
}

func main() {
	file, err := os.Open("movements.log")
	if err != nil {
		log.Fatal(err)
	}

	incoming, err := os.Create("processed/incoming.log")
	if err != nil {
		log.Fatal(err)
	}
	outgoing, err := os.Create("processed/outgoing.log")
	if err != nil {
		log.Fatal(err)
	}

	inError, err := os.Create("processed/inError.log")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	balance := 0
	linecount := 0

	for scanner.Scan() {
		t := Transaction{}

		t.populateFromLog(scanner.Text())

		if !t.isValid {
			// If invalid, print not the Transaction type but the unmodified line so it can be corrected later.
			_, err := inError.Write([]byte(fmt.Sprintf("%s\n", []byte(scanner.Text()))))
			if err != nil {
				log.Fatal(err)
			}
			linecount++
			continue
		}

		if t.Type == "OUT" {
			if err := t.writeToLog(outgoing); err != nil {
				log.Fatal(err)
			}
			balance -= t.Amount
		} else if t.Type == "IN" {
			if err := t.writeToLog(incoming); err != nil {
				log.Fatal(err)
			}
			balance += t.Amount
		}

		linecount++
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	fmt.Printf("Count:\t\t%d\nBalance:\t%d\n", linecount, balance)
}
