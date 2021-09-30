package main

import (
	"fmt"
	"log"
	"io/ioutil"
)

func main() { // Logging is threat safe especially in overlapping, fmt is not
	log.Println("Log entry")

	log.SetFlags(0)

	for i := 0; i < 100; i+=1 {
		go log.Println(i)
	}

	for i := 0; i < 100; i+=1 {
		go fmt.Println(i)
	}

	log.SetOutput(ioutil.Discard)

	log.Println("Entry 2")

	defer log.Println("Will not be logged")
	log.Fatal("Exit")
}