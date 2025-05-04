package main

import (
	"fmt"
)

func main() {
	// 1. 字面量初始化
	s1 := []int{1, 2, 3}
	fmt.Println("1. 字面量:", s1)

	// 2. 使用 make 创建切片
	s2 := make([]int, 3)    // 长度=3, 容量=3，默认值全是0
	s3 := make([]int, 2, 5) // 长度=2, 容量=5
	fmt.Println("2. make([]int, 3):", s2)
	fmt.Println("   make([]int, 2, 5):", s3)

	// 3. 空切片和 nil 切片
	var s4 []int  // nil 切片，尚未分配内存
	s5 := []int{} // 空切片，但已分配内存
	fmt.Println("3. nil 切片s4:", s4, "is nil?", s4 == nil)
	fmt.Println("   空切片s5:", s5, "is nil?", s5 == nil)

	// 4. 从数组切片
	arr := [5]int{10, 20, 30, 40, 50}
	s6 := arr[1:4] // 切片视图 [20 30 40]
	fmt.Println("4. 数组切片:", s6)

	// 5. 从 nil 切片 append
	var s7 []int
	s7 = append(s7, 100, 200, 300)
	fmt.Println("5. append 到 nil 切片:", s7)

	// 6. 从函数返回切片
	s8 := getSlice()
	fmt.Println("6. 从函数返回的切片:", s8)

	// 7. 组合初始化（复制数据）
	original := []int{7, 8, 9}
	s9 := append([]int{}, original...) // 创建新切片，避免共享底层数组
	fmt.Println("7. 组合初始化:", s9)

	// 8. 切片操作：增、删、改、查
	ops := []int{1, 2, 3, 4, 5}
	fmt.Println("原始切片:", ops)

	// 查：访问第三个元素
	fmt.Println("查 - 第三个元素:", ops[2])

	// 改：将第二个元素改为20
	ops[1] = 20
	fmt.Println("改 - 第二个元素改为20:", ops)

	// 增：在末尾添加元素6,7
	ops = append(ops, 6, 7)
	fmt.Println("增 - 添加元素6,7:", ops)

	// 删：删除下标为2的元素（值为3）
	ops = append(ops[:2], ops[3:]...)
	fmt.Println("删 - 删除下标为2的元素:", ops)
	// 9. 遍历切片
	fmt.Println("9. 遍历切片:")
	for i := 0; i < len(ops); i++ {
		fmt.Printf("index %d: value %d\n", i, ops[i])
	}

	for idx, val := range ops {
		fmt.Printf("range - index %d: value %d\n", idx, val)
	}
}

func getSlice() []int {
	return []int{400, 500, 600}
}
