package main

import "fmt"

type Cat struct {
	name    string
	age     int
	variety string
}

func (cat *Cat) getName() string {
	return cat.name
}

func (cat *Cat) setName(name string) {
	cat.name = name
}

func (cat *Cat) show() {
	fmt.Printf("名字：%s, 年龄：%d, 品种: %s", cat.name, cat.age, cat.variety)
}

func (cat *Cat) eat() {
	fmt.Printf("%s在吃鱼！\n", cat.name)
}

func (cat *Cat) sleep() {
	fmt.Printf("%s在睡觉!~ \n", cat.name)
}

func newCat() Cat {
	return Cat{
		name:    "小蓝",
		age:     1,
		variety: "橘猫",
	}
}

func main() {
	cat := newCat()

	fmt.Printf("你好我叫%s \n", cat.getName())
	cat.setName("小小")
	fmt.Printf("我改名啦，我现在叫%s \n", cat.getName())
	cat.sleep()
	cat.eat()
	cat.show()
}
