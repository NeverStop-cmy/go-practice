package functions

import "fmt"

func RunBasicFunctionsDemo() {
	fmt.Println("\n基本函数演示:")

	// 基本函数
	sum := add(3, 5)
	fmt.Println("3 + 5 =", sum)

	// 多返回值
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}

func add(a, b int) int {
	return a + b
}

func swap(x, y string) (string, string) {
	return y, x
}
