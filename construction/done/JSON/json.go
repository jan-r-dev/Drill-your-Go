package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type People struct {
	People []*Person `json:"people,omitempty"`
}

func (ppl *People) readPeople() {
	fmt.Printf("Reading file now...\n\n")

	for i, v := range ppl.People {
		fmt.Printf("Index: %d\n\tFirst name: %q\n\tLast name: %q\n\tAge: %d\n", i, v.FirstName, v.LastName, v.Age)
	}
}

func (ppl *People) markPeopleAsProcessed() {
	fmt.Printf("\nModifying file now...\n\n")

	for _, v := range ppl.People {
		v.Processed = true
	}
}

type Person struct {
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	Gender    string `json:"gender,omitempty"`
	Age       int    `json:"age,omitempty"`
	Number    string `json:"number,omitempty"`
	Processed bool   `json:"processed,omitempty"`
}

func main() {
	ppl := &People{}
	loadJSON(ppl, "sample.json")

	ppl.readPeople()

	ppl.markPeopleAsProcessed()

	saveJson(ppl, "processed.json")
}

func loadJSON(ppl *People, filename string) {
	fmt.Printf("Ingesting file now...\n\n")
	bs, err := ioutil.ReadFile("sample.json")
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(bs, ppl)
	if err != nil {
		log.Fatal(err)
	}
}

func saveJson(ppl *People, filename string) {
	bs, err := json.MarshalIndent(ppl, "", "    ")
	if err != nil {
		log.Fatal(err)
	}

	os.WriteFile(filename, bs, 0644)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("File %q written successfully.\n", filename)
	}
}
