package main

import (
	"log"
	l "log" // only useful if we’ve package which has the same interface (exported identifiers) as other imported package

	_ "log"  // will only call init if any
	. "math" // will break if 2 packages have common exported identifiers
)

func main() {
	log.Println("Log Entry: 1")
	l.Println("Log Entry: 2")
	l.Printf("Log Entry: %d", int(Floor(3.3)))
}
