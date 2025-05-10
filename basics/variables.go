package basics

import "fmt"

func RunVariablesDemo() {
	fmt.Println("\n变量和常量演示:")

	// 变量声明
	var name string = "Go"
	age := 30 // 类型推断

	// 常量
	const pi = 3.14159

	fmt.Println("Name:", name, "Age:", age, "Pi:", pi)
}
