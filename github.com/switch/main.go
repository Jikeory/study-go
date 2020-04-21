package main

import (
	"fmt"
)

func main() {
	var x interface{}
	var y = 10.0
	x = y
	switch i := x.(type) {
	case nil :
		fmt.Printf("%T\n", i)
	case float64 :
		fmt.Printf("%T\n", i)
	case int :
		fmt.Printf("%T\n", i)
	case bool, string :
		fmt.Printf("%T\n", i)
	default:
		fmt.Println("interface")
	}
}
