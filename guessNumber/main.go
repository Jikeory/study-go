package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	num := rand.Intn(99) + 1
	fmt.Printf("num=%v\n", num)
	for i := 1; i <= 10; i++ {
		var number int
		fmt.Println("请输入猜测的1-100的数字")
		fmt.Scan(&number)
		if i == 1 {
			if num == number {
				fmt.Println("你真是个天才")
				break
			} else {
				continue
			}
		} else if i <= 3 {
			if num == number {
				fmt.Println("你很聪明")
				break
			} else {
				continue
			}
		} else if i <= 9 {
			if num == number {
				fmt.Println("一般般")
				break
			} else {
				continue
			}
		} else if i == 10 {
			if num == number {
				fmt.Println("可算猜对了")
				break
			} else {
				continue
			}
		} else {
			fmt.Println("猜错了")
			break
		}
	}
}
