package main

import (
	"fmt"
)

func main() {
	var name = "yuyang"
	fmt.Println(name)

	var input string
	fmt.Print("在次输入：")
	fmt.Scanln(&input)
	fmt.Print(input)

	fmt.Scan(&input)
	fmt.Println(input)
}
