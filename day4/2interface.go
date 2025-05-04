package main

import (
	"fmt"
	"math"
)

// 定义形状接口
type Shape interface {
	Area() float64
	Perimeter() float64
}

// 实现接口的结构体：矩形
type Rectangle struct {
	Width  float64
	Height float64
}

// 实现接口的结构体：圆形
type Circle struct {
	Radius float64
}

// 矩形实现Shape接口的方法
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// 圆形实现Shape接口的方法
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// 创建形状的工厂函数
func NewShape(shapeType string, params ...float64) Shape {
	switch shapeType {
	case "rectangle":
		if len(params) >= 2 {
			return Rectangle{Width: params[0], Height: params[1]}
		}
	case "circle":
		if len(params) >= 1 {
			return Circle{Radius: params[0]}
		}
	}
	return nil
}

// 打印形状信息
func printShapeInfo(s Shape) {
	fmt.Printf("面积: %.2f\n", s.Area())
	fmt.Printf("周长: %.2f\n", s.Perimeter())
}

func main() {
	// 创建并使用矩形
	rect := NewShape("rectangle", 5.0, 3.0)
	fmt.Println("矩形:")
	printShapeInfo(rect)

	// 创建并使用圆形
	circle := NewShape("circle", 2.0)
	fmt.Println("\n圆形:")
	printShapeInfo(circle)

	// 接口切片
	shapes := []Shape{rect, circle}

	// 遍历所有形状
	for i, shape := range shapes {
		fmt.Printf("\n形状 %d:\n", i+1)
		printShapeInfo(shape)
	}
}
