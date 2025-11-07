package main

import (
	"errors"
	"fmt"
	"time"
)

const (
	TimeFmt = "2006-01-02 15:04:05"
)

type BusinessError struct {
	Code    int
	Message string
	Time    time.Time
}

func (e *BusinessError) Error() string {
	return fmt.Sprintf("错误代码：%d, 消息：%s, 时间：%s",
		e.Code, e.Message, e.Time.Format(TimeFmt))
}

func queryDB(userID int) (string, error) {
	if userID <= 0 {
		return "", &BusinessError{
			Code:    1001,
			Message: "数据不存在",
			Time:    time.Now(),
		}
	}
	//模拟查询
	if userID == 999 {
		//返回用户信息
		return "", errors.New("数据库连接超时")
	}
	return "用户数据", nil

}

func main() {
	//创建模拟查询实例
	userinfo, err := queryDB(-1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(userinfo)
	}

}
