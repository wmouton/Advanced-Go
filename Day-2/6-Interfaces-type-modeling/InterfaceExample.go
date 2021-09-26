package main

import (
	"fmt"
)

func main() {
	var h Helper = HelpString("Help me")
	fmt.Println(h.Help())

	var explicit = interface {Help() string}.Help(h)
	fmt.Println(explicit)

	// Polymorphism
	var helpers = []Helper {
		HelpString("Help me again"),
		&UnHelpString{},
	}

	fmt.Println(helpers)

	for _, hleper := range helpers {
		fmt.Println(helper.Help())
	}

	var h2 interface{} = HelpString("Please help me more")
	n, ok := h2.(Helper)
	fmt.Println(n, ok)
	
	var _ = h2h.(string) // panic
}