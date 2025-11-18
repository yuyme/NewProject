package main

import (
	"fmt"
	"time"
)

const (
	UserTypeNormal = "normal"
	UserTypeVIP    = "vip"
)

type BaseUser struct {
	ID        int64
	UserName  string
	Email     string
	CreatedAt time.Time
}

func (u *BaseUser) GetCreateDate() string {
	return u.CreatedAt.Format(time.DateTime)
}

func (u *BaseUser) DisplayBasicInfo() {
	fmt.Printf("用户ID：%d，用户名：%s，邮箱：%s，创建时间：%s\n", u.ID, u.UserName, u.Email, u.GetCreateDate())
}

type Address struct {
	Province string
	City     string
	Street   string
	//详细地址
	Detail string
}

func (a Address) GetFullAddress() string {
	return fmt.Sprintf("%s%s%s%s", a.Province, a.City, a.Street, a.Detail)
}

type NormalUser struct {
	BaseUser
	Address []Address
}

func (nu *NormalUser) AddAddress(address Address) {
	nu.Address = append(nu.Address, address)
}

type VIPUser struct {
	BaseUser
	Address    []Address
	VIPLevel   int
	Discount   float64
	ExpireTime time.Time
}

func (vu *VIPUser) AddAddress(address Address) {
	vu.Address = append(vu.Address, address)
}

func (vu *VIPUser) IsVipValid() bool {
	return vu.ExpireTime.After(time.Now())
}

func (vu *VIPUser) GetDiscount() float64 {
	if vu.IsVipValid() {
		return vu.Discount
	}
	return 1.0
}

func genUserID() int64 {
	return time.Now().Unix()
}

type UserService struct {
}

func (us *UserService) CreatUser(username, email string, userType string) (interface{}, error) {
	base := BaseUser{
		ID:        genUserID(),
		UserName:  username,
		Email:     email,
		CreatedAt: time.Now(),
	}

	switch userType {
	case UserTypeNormal:
		return NormalUser{
			BaseUser: base,
			Address:  []Address{},
		}, nil
	case UserTypeVIP:
		vip := VIPUser{
			BaseUser:   base,
			Address:    []Address{},
			VIPLevel:   1,
			Discount:   0.9,
			ExpireTime: time.Now().AddDate(0, 0, 30),
		}
		return vip, nil
	default:
		return nil, fmt.Errorf("无效的用户类型！")
	}
}

func main() {
	/*//初始化一个用户
	service := &UserService{}
	normalUser, err := service.CreatUser("张三", "zhangsan@example.com", UserTypeNormal)
	if err != nil {
		fmt.Println(UserTypeNormal, err)
		return
	}

	vipUser, err := service.CreatUser("李四", "lisi@example.com", UserTypeVIP)
	if err != nil {
		fmt.Println(UserTypeVIP, err)
		return
	}
	nUser := normalUser.(NormalUser)
	nUser.AddAddress(Address{
		Province: "河南省",
		City:     "平顶山市",
		Street:   "叶县",
		Detail:   "九龙街道",
	})
	nUser.DisplayBasicInfo()

	//多态是实列
	users := []interface{}{nUser, vipUser}
	//vUser := vipUser.(VIPUser)
	for _, user := range users {
		switch u := user.(type) {
		case NormalUser:
			fmt.Println("普通用户")
			fmt.Println("用户id：", u.ID)
			fmt.Println("用户名：", u.UserName)
			fmt.Println("邮箱：", u.Email)
			fmt.Println("创建时间：", u.GetCreateDate())
			fmt.Print("地址：")
			for _, add := range u.Address {
				fmt.Println("-", add.GetFullAddress())
			}
		case VIPUser:
			fmt.Println("VIP用户")
			fmt.Println("用户id：", u.ID)
			fmt.Println("用户名：", u.UserName)
			fmt.Println("邮箱：", u.Email)
			fmt.Println("创建时间：", u.GetCreateDate())
			fmt.Println("vip等级：", u.VIPLevel)
			fmt.Println("vip状态：", u.IsVipValid())
			fmt.Println("折扣：", u.GetDiscount())
			fmt.Println("vip有效期：", u.ExpireTime)
			fmt.Print("地址：")
			for _, add := range u.Address {
				fmt.Println("-", add.GetFullAddress())
			}
		}
	}*/
	notifierMain()

}
