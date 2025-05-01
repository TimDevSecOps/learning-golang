package main

import (
	"errors"
	"fmt"
)

func add(x, y int) int {
	return x + y
}

func divide(x, y float64) (float64, error) {
	if y == 0 {
		return 0, errors.New("division by zero")
	}
	return x / y, nil
}

// 命名返回值
func rectangle(width, height float64) (area, perimeter float64) {
	area = width * height
	perimeter = 2 * (width + height)
	return // 隐式返回命名的返回值
}

// 可变参数
func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

// 函数作为参数
func calculate(x, y int, operation func(int, int) int) int {
	return operation(x, y)
}

// 闭包
func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func main() {
	// 基本函数调用
	fmt.Println("1 + 2 =", add(1, 2))

	// 多返回值处理
	if result, err := divide(10, 2); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("10 / 2 =", result)
	}

	// 命名返回值
	area, perim := rectangle(3, 4)
	fmt.Printf("Rectangle area: %.2f, perimeter: %.2f\n", area, perim)

	// 可变参数
	fmt.Println("sum:", sum(1, 2, 3, 4, 5))

	// 函数作为参数
	multiply := func(x, y int) int { return x * y }
	fmt.Println("3 * 4 =", calculate(3, 4, multiply))

	// 闭包
	counter := counter()
	fmt.Println("Counter:", counter())
	fmt.Println("Counter:", counter())
	fmt.Println("Counter:", counter())
}

/*
不支持的情况：
1. 函数重载（Go不支持函数重载）
func add(x, y int) int { return x + y }
func add(x, y, z int) int { return x + y + z } // 编译错误

2. 默认参数值
func greet(name string = "World") { } // 编译错误

3. 函数嵌套定义（只支持匿名函数）
func outer() {
    func inner() { } // 编译错误
}
*/
