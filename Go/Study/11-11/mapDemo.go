package main

import (
	"fmt"
	"sync"
)

func main() {
	//模拟set集合 (无序 不重复)
	uniqueIDs := make(map[int]bool)
	ids := []int{1, 2, 2, 3, 4, 5}

	for _, id := range ids {
		uniqueIDs[id] = true
	}

	fmt.Printf("数组IDs{%v}的大小为：%v \n", uniqueIDs, len(uniqueIDs))

	//键值反转
	originMap := map[string]int{
		"one":   1,
		"tow":   2,
		"three": 3,
	}
	reversedMap := make(map[int]string)
	for key, value := range originMap {
		reversedMap[value] = key
	}

	fmt.Printf("键值反转前originMap：%v \n", originMap)
	fmt.Printf("键值反转后reversedMap：%v \n", reversedMap)

	//另一种map，并发
	var syncMap sync.Map
	//添加值
	syncMap.Store("key1", "value1")
	syncMap.Store("key2", "value2")
	//取值
	fmt.Println(syncMap.Load("key1"))
	//移除元素
	syncMap.Delete("key2")
	//遍历
	syncMap.Range(func(key, value interface{}) bool {
		fmt.Printf("键：%v , 值：%v", key, value)
		return true
	})
}
