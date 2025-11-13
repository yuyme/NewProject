package main

import (
	"fmt"
	"sort"
	"time"
)

type Product struct {
	ID        int
	Name      string
	Price     float64
	Rating    float64
	Sales     int
	CreatedAt time.Time
}

type ProductManger struct {
	products []Product
}

func NewProductManger() *ProductManger {
	return &ProductManger{
		products: []Product{
			{1, "无线耳机", 299.99, 4.5, 1200, time.Now().Add(-24 * time.Hour)},
			{2, "只能手机", 3999.99, 4.8, 850, time.Now().Add(-48 * time.Hour)},
			{3, "华强北手表", 99.99, 4.2, 2300, time.Now().Add(-12 * time.Hour)},
			{4, "蓝牙无线耳机", 199.99, 4.6, 1500, time.Now().Add(-72 * time.Hour)},
		},
	}
}

func (pm *ProductManger) SortByPrice(ascending bool) {
	if ascending {
		sort.Slice(pm.products, func(i, j int) bool {
			return pm.products[i].Price < pm.products[j].Price
		})
	} else {
		sort.Slice(pm.products, func(i, j int) bool {
			return pm.products[i].Price > pm.products[j].Price
		})
	}

}

func (pm *ProductManger) showProductMan() {
	for _, p := range pm.products {
		fmt.Printf("ID：%v Name：%v Price：%v Rating：%v Sales：%v \n",
			p.ID, p.Name, p.Price, p.Rating, p.Sales)
	}
}

func main() {
	pm := NewProductManger()

	fmt.Println("=================按价格排序==================")
	pm.SortByPrice(false)
	pm.showProductMan()
}
