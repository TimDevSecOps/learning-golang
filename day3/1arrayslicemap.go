package main

import "fmt"

func main() {
	// 数组：固定长度
	var arr [3]int = [3]int{1, 2, 3}
	fmt.Printf("%%v: %v\n", arr)   // %v: [1 2 3]
	fmt.Printf("%%#v: %#v\n", arr) // %#v: [3]int{1, 2, 3}
	fmt.Printf("%%T: %T\n", arr)   // %T: [3]int

	var t bool = true
	fmt.Printf("%%t: %t\n", t)        // %t: true
	fmt.Printf("%%b: %b\n", arr[2])   // %b: 11 (b11 == 3)
	fmt.Printf("%%+b: %+b\n", arr[2]) // %+b: +11 (b11 == 3)
	fmt.Printf("%%#b: %#b\n", arr[2]) // %#b: 0b11 (b11 == 3)

	var s string = "hello"
	fmt.Printf("%%c: %c\n", s[0]) // %c: h
	fmt.Printf("%%q: %q\n", s)    // %q: "hello"

	fmt.Printf("output: %[3]*.[2]*[1]f\n", 12.036, 2, 6) // 6.2f: output: 12.04
	fmt.Printf("%d %d %#[1]x %#x\n", 16, 17)             // 16 17 0x10 0x11

	fmt.Printf("数组: %v, 长度: %d, 类型: %T\n", arr, len(arr), arr)

	arr[0] = 100
	fmt.Printf("数组: %v, 长度: %d\n", arr, len(arr))

	// 切片：动态长度
	slice := make([]int, 3, 5) // 长度3，容量5
	slice = append(slice, 4)   // 动态添加元素
	fmt.Printf("切片: %v, 长度: %d, 容量: %d\n", slice, len(slice), cap(slice))

	// 切片操作
	numbers := []int{1, 2, 3, 4, 5}
	sub := numbers[1:3] // [2 3]
	fmt.Printf("子切片: %v, %T\n", sub, sub)

	// 映射
	ages := map[string]int{
		"Alice": 25,
		"Bob":   30,
	}

	// 添加和删除
	ages["Charlie"] = 35
	delete(ages, "Bob")

	// 安全访问
	if age, exists := ages["Alice"]; exists {
		fmt.Printf("Alice的年龄: %d\n", age)
	}
}

/*
不支持的情况：
1. 数组不支持负数索引
arr[-1] = 1 // 编译错误

2. 映射键类型限制
type Person struct{}
m := map[Person]string{} // 如果Person不能比较，则编译错误

3. 切片容量不能小于长度
s := make([]int, 5, 3) // 编译错误

4. 数组不能改变大小
var arr [3]int
arr[3] = 4 // 编译错误：越界

5. 映射不保证遍历顺序
for k, v := range map{} // 顺序是随机的
*/
