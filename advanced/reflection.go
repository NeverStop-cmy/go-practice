package advanced

import (
	"fmt"
	"reflect"
)

func RunReflectionDemo() {
	var x float64 = 3.14

	// 获取类型信息
	t := reflect.TypeOf(x)
	fmt.Println("Type:", t)

	// 获取值信息
	v := reflect.ValueOf(x)
	fmt.Println("Value:", v)

	// 修改值(必须是指针)
	p := reflect.ValueOf(&x).Elem()
	p.SetFloat(6.28)
	fmt.Println("New value:", x)
}
