package main

import "fmt"

/*func validateUser(username, password string) (bool, error) {
	if len(username) < 3 {
		return false, fmt.Errorf("用户名不能小于3位")
	}
	if len(password) < 8 {
		return false, fmt.Errorf("密码不能少于8位")
	}
	return true, nil
}*/

func creatCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func main() {
	counter := creatCounter()
	for _ = range 5 {
		fmt.Println("creatCounter：", counter())
	}
}

/*username := "zhangsan"
password := "1234567"

flag, err := validateUser(username, password)

if flag {
	fmt.Println("登陆成功")
} else {
	fmt.Println(err)
}*/
