package oop

import "fmt"

type Person struct {
	Name string
	Age  int
}

func RunStructsDemo() {
	// 创建结构体实例
	p1 := Person{Name: "Alice", Age: 25}
	p2 := Person{"Bob", 30}

	fmt.Println(p1, p2)

	// 使用指针访问结构体
	p3 := &Person{"Charlie", 35}
	fmt.Println((*p3).Name) // 可以省略*
	fmt.Println(p3.Name)    // 等价于上面
}
