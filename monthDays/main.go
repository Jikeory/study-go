package main

import "fmt"

func days() {
	for {
		var year int
		var month int
		var day int
		fmt.Println("请输入年份。。。")
		fmt.Scan(&year)
		if year > 2020 || year < 0 {
			fmt.Println("输入的年份有误")
			continue
		}
		fmt.Println("请输入月份。。。")
		fmt.Scan(&month)
		if month > 12 || month < 1 {
			fmt.Println("输入的月份有误")
			continue
		}
		fmt.Println("请输入日期。。。")
		fmt.Scan(&day)
		switch month {
		case 1, 3, 5, 7, 8, 10, 12:
			if day > 31 || day < 1 {
				fmt.Println("输入的日期有误")
				continue
			}
		case 2:
			if year % 4 == 0 {
				if day > 30 || day < 1 {
					fmt.Println("输入的日期有误")
					continue
				}
			}
		default:
			if day > 30 || day < 1 {
				fmt.Println("输入的日期有误")
				continue
			}
		}
		fmt.Printf("输入的日期是%v年%v月%v", year, month, day)
		break
	}
}

func main() {
	//fmt.Println(1)
	days()
}
