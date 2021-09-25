package main

import "fmt"

func inbetween(from, to int) []int {
    var rslt []int
    
    if from > to {
        rslt = []int{}
    } else {
        for i := from; i < to; i++ {
            rslt = append(result, i)
        }
    }
    
    return rslt
}

func main() {
    for i := range inbetween(0, 10) {
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
