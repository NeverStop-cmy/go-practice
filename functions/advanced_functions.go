package functions

import "fmt"

// 命名返回值
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return // 裸返回
}

func deferDemo() {
	defer fmt.Println("世界") // defer语句会延迟执行
	fmt.Println("你好")
}

func RunAdvancedFunctionsDemo() {
	fmt.Println("\n高级函数演示:")
	fmt.Println(split(17))
	deferDemo()
}
