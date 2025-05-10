package basics

import "fmt"

func RunDataTypesDemo() {
	fmt.Println("\n基本数据类型演示:")
	// 整数
	var a int = 42
	var b uint = 10

	// 浮点数
	var c float64 = 3.14

	// 布尔值
	var isGoCool bool = true

	// 字符串
	var greeting string = "Hello, Go!"

	fmt.Println(a, b, c, isGoCool, greeting)
}
