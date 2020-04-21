package main

import "fmt"

var name string = "jack"

func test1() {
	fmt.Println(name)
}

func test2() {
	name = "tom"
	fmt.Println(name)
}

func main() {
	fmt.Printf("name=%v", name)
	test1()
	test2()
	fmt.Printf("name=%v", name)
	test1()
}