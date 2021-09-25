package main

import "fmt"

func boober(from, to int) []int {
	var resulting []int

	if from > to {
		resulting = []int{}
	} else {
		for i := from; i < to; i++ {
			resulting = append(resulting, i)
		}
	}

	return resulting
}

func main() {
	for i := range boober(0, 10) {
		switch i % 5 {
		case 1:
			fmt.Println("fizz")
		case 2:
			fmt.Println("bazz")
		case 3:
			fmt.Println("gizz")
			fallthrough
		default:
			fmt.Println("fizzbazz")
		}
	}
}