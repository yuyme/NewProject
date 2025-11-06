package main

import (
	"fmt"
)

func main() {
	price := 99

	fmt.Println(price)

	//成绩判断
	//fmt.Print("请输入你的成绩：")
	//var score int
	//fmt.Scan(&score)
	score := 72

	if score >= 90 {
		fmt.Println("优秀")
	} else if score >= 80 {
		fmt.Println("良好")
	} else if score >= 70 {
		fmt.Println("一般")
	} else if score >= 60 {
		fmt.Println("及格")
	} else {
		fmt.Println("较差 不及格")
	}

	score2 := 92

	switch {
	case score2 >= 90:
		fmt.Println("优秀")
	case score2 >= 80:
		fmt.Println("良好")
	case score2 >= 70:
		fmt.Println("一般")
	case score2 >= 60:
		fmt.Println("及格")
	default:
		fmt.Println("较差 不及格")
	}

	count := 0
	for count >= 0 {
		fmt.Println("hello go")
		count++
	}

}
