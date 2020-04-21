package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Printf()

	var str1 string
	var str2 string
	var str3 string
	var i int = 100
	var f float64 = 12.13
	var b bool = true

	str1 = strconv.FormatInt(int64(i), 10)
	fmt.Printf("%T %v\n", str1, str1)
	str2 = strconv.FormatFloat(12.13, 'f', 2, 64)
	fmt.Printf("%T %v\n", str2, str2)
	str3 = strconv.FormatBool(b)
	fmt.Printf("%T %v\n", str3, str3)
}
