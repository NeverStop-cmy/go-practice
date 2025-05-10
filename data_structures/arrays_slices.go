package data_structures

import "fmt"

func RunArraysSlicesDemo() {
	fmt.Println("\n数组、切片类型演示:")
	// 数组
	var arr [5]int
	arr[0] = 1
	arr[1] = 2
	fmt.Println(arr)

	// 切片(动态数组)
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)

	// 切片操作
	subSlice := slice[1:4]
	fmt.Println(subSlice)

	// 使用make创建切片
	s := make([]int, 5, 10) // 长度5，容量10
	fmt.Println(len(s), cap(s))

	//切片的遍历
	ca := make([]int, 0, 10) // 长度0，容量10
	ca = append(ca, 1)
	for index, value := range ca {
		fmt.Println(index, value)
	}
}
