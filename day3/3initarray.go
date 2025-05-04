package main

import "fmt"

func main() {
	// 1. 完整初始化：指定所有元素
	var a1 = [3]int{1, 2, 3}
	fmt.Println("1. 完整初始化:", a1)

	// 2. 部分初始化：未指定的元素自动为零值
	var a2 = [5]int{1, 2}
	fmt.Println("2. 部分初始化:", a2)

	// 3. 自动推导长度
	a3 := [...]int{10, 20, 30}
	fmt.Println("3. 自动推导长度:", a3)

	// 4. 指定下标初始化（稀疏数组）
	a4 := [5]int{1: 100, 3: 300}
	fmt.Println("4. 指定下标初始化:", a4)

	// 5. 字符数组（可作为字符串处理）
	var a5 = [6]byte{'H', 'e', 'l', 'l', 'o', '!'}
	fmt.Println("5. 字符数组（打印为字符）:", string(a5[:])) // string(a5) 会报错，不能将数组转换为字符串

	// 6. 多维数组
	var a6 = [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println("6. 多维数组:", a6)

	// 7. 默认零值数组
	var a7 [4]string
	fmt.Println("7. 默认零值数组:", a7)

	// ===== 数组操作示例 =====

	// 查：访问数组元素
	fmt.Println("查: a1[1] =", a1[1]) // 访问第二个元素

	// 改：修改数组元素
	a1[1] = 20
	fmt.Println("改: 修改后 a1 =", a1)

	// 增：由于数组是固定长度，不能直接“增加”元素，但可以构造新数组模拟
	a1New := [4]int{a1[0], a1[1], a1[2], 99}
	fmt.Println("增: 新数组 a1New =", a1New)

	// 删：不能直接删除元素，但可以通过构造新数组忽略指定元素模拟删除
	// 删除 a1 中第 1 个元素（索引为 0）
	a1Del := [2]int{a1[1], a1[2]}
	fmt.Println("删: 删除索引 0 后 a1Del =", a1Del)
	// 遍历：使用 for 循环或 range 遍历数组
	fmt.Println("遍历 a1:")
	for i := 0; i < len(a1); i++ {
		fmt.Printf("索引 %d -> 值 %d\n", i, a1[i])
	}

	fmt.Println("使用 range 遍历 a3:")
	for idx, val := range a3 {
		fmt.Printf("a3[%d] = %d\n", idx, val)
	}
}
