package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type Details struct {
	words  int
	spaces int
	lines  int
}

func main() {
	//Reading file
	content, err := ioutil.ReadFile("file.txt")
	if err != nil {
		log.Fatal(err)
	}

	//Assigned values to structure
	para := Details{
		lines:  strings.Count(string(content), "."),
		spaces: strings.Count(string(content), " "),
		words:  len(strings.Fields(string(content))),
	}

	value := &para //value will hold the updated details of structure

	fmt.Println(string(content))
	fmt.Println("Number Of Words : ", value.words)
	fmt.Println("Number Of Spaces : ", value.spaces)
	fmt.Println("Number Of Lines : ", value.lines)
}
