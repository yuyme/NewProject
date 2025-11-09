package main

import (
	"fmt"
)

/*
 数组和切片
 数组声明方式
 var arr1 [3]int                    // 声明但不初始化
 arr2 := [3]int{1, 2, 3}           // 声明并初始化
 arr3 := [...]int{1, 2, 3, 4, 5}   // 编译器计算长度
 切片是对数组的抽象，提供动态大小的灵活视图。切片是引用类型，包含三
 切片创建方式
 var slice1 []int                    // 声明切片
 slice2 := make([]int, 3, 5)        // 使用make，长度3，容量5
 slice3 := []int{1, 2, 3}           // 字面量声明
 slice4 := arr2[1:3]                // 从数组创建
*/

// 模拟用户购物车行为
// 创建一个购物车对象
type shopCar struct {
	Name     string
	Price    float64
	Quantity int
}

func main() {
	//创建一个购物车对象数组		这里可以用切片
	var shopCars []shopCar

	//模拟往购物车里面添加商品
	shopCars = append(shopCars, shopCar{
		Name:     "XIAOMI 17 Pro Max",
		Price:    5999.00,
		Quantity: 9,
	})

	shopCars = append(shopCars, shopCar{
		Name:     "Mac Book Pro",
		Price:    19999.99,
		Quantity: 1,
	})

	shopCars = append(shopCars, shopCar{
		Name:     "XIAOMI Watch S3",
		Price:    1099.00,
		Quantity: 100,
	})

	total := 0
	monTotal := 0.0

	//遍历购物车获取总金额和商品数量
	for s := range shopCars {
		fmt.Printf("商品名: %v, 商品价格：%v, 商品数量：%v \n", shopCars[s].Name, shopCars[s].Price, shopCars[s].Quantity)
		total += shopCars[s].Quantity
		monTotal += shopCars[s].Price
	}

	fmt.Printf("商品合计金额为：%v \n", monTotal)
	fmt.Printf("商品总数为：%v", total)

}
