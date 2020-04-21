package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str string = "true"
	var b bool
	var i1 int64
	var i2 int32
	var f1 float64
	var f2 float32

	b , _ = strconv.ParseBool(str)
	fmt.Printf("%T %v\n", b, b)

	i1 , _ = strconv.ParseInt(str, 10, 64)
	i2 = int32(i1)
	fmt.Printf("%T %v\n", i2, i2)
	f1 , _ = strconv.ParseFloat(str, 64)
	f2 = float32(f1)
	fmt.Printf("%T %v\n", f2, f2)

}
