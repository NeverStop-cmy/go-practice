package oop

import (
	"fmt"
	"math"
)

type Rectangle struct {
	Width, Height float64
}

// 为Rectangle类型定义方法
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r *Rectangle) Scale(factor float64) {
	r.Width *= factor
	r.Height *= factor
}

// 定义接口
type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	Radius float64
}

type Square struct {
	Side float64
}

// Circle的方法实现
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Square的方法实现
func (s Square) Area() float64 {
	return s.Side * s.Side
}

func (s Square) Perimeter() float64 {
	return 4 * s.Side
}

// 接受接口类型的函数
func PrintInfo(s Shape) {
	fmt.Printf("Area: %.2f, Perimeter: %.2f\n", s.Area(), s.Perimeter())
}

func RunMethodsInterfacesDemo() {
	rect := Rectangle{Width: 10, Height: 5}
	fmt.Println("Area:", rect.Area())

	rect.Scale(2)
	fmt.Println("Scaled Area:", rect.Area())

	c := Circle{Radius: 5}
	s := Square{Side: 4}

	PrintInfo(c)
	PrintInfo(s)

	// 接口数组
	shapes := []Shape{c, s}
	for _, shape := range shapes {
		PrintInfo(shape)
	}

}
