package main

import (
	"fmt"
	_ "strconv"
)

func main() {
	var i int8 = 10
	// ptr存放的是i的值10的内存地址
	var ptr *int8 = &i
	fmt.Printf("%v\n", ptr)
	fmt.Printf("%v", *ptr)
}
