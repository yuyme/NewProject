package main

import "fmt"

/*
func validateUser(username, password string) (bool, error) {
	if len(username) < 3 {
		return false, fmt.Errorf("用户名不能小于3位")
	}
	if len(password) < 8 {
		return false, fmt.Errorf("密码不能少于8位")
	}
	return true, nil
}*/

/*函数闭包
func creatCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}*/

/*
创建类
*/
type User struct {
	ID      int
	Name    string
	Balance float32
}

func updateBalanceValue(user User, amount float32) {
	user.Balance += amount
	fmt.Println("updateBalanceValue 更新后的余额：", user.Balance)
}

func updateBalancePointer(user *User, amount float32) {
	user.Balance += amount
	fmt.Println("updateBalancePointer 更新后的余额：", user.Balance)
}

func main() {
	user := User{
		ID:      1,
		Name:    "zhangsan",
		Balance: 100.00,
	}
	//值传递
	updateBalanceValue(user, 50)
	fmt.Println("updateBalanceValue", user)
	//指针传递 (内存地址传递)
	updateBalancePointer(&user, 60)
	fmt.Println("updateBalancePointer", user)
	fmt.Println("---------------分割线-----------------")
	users := []*User{
		&User{
			ID:      2,
			Name:    "lisi",
			Balance: 3000.0,
		},
		&User{
			ID:      3,
			Name:    "wangwu",
			Balance: 10000.0,
		},
		&User{
			ID:      4,
			Name:    "zhaoliu",
			Balance: 4564.21,
		},
	}
	for _, u := range users {
		fmt.Println("发工资之前的余额：", u)
	}

	for _, u := range users {
		u.Balance += 3500
	}
	fmt.Println(".................................")

	for _, u := range users {
		fmt.Println("发完工资后的余额：", u)
	}
	/*
		函数闭包counter := creatCounter()
		for _ = range 5 {
			fmt.Println("creatCounter：", counter())
		}*/
}

/*if校验
username := "zhangsan"
password := "1234567"

flag, err := validateUser(username, password)

if flag {
	fmt.Println("登陆成功")
} else {
	fmt.Println(err)
}*/
