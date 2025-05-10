package data_structures

import "fmt"

func RunMapsDemo() {
	// 创建map
	m := make(map[string]int)

	// 添加元素
	m["one"] = 1
	m["two"] = 2

	// 获取元素
	fmt.Println(m["one"])

	// 删除元素
	delete(m, "two")

	// 检查键是否存在
	value, exists := m["two"]
	fmt.Println(value, exists)
}
