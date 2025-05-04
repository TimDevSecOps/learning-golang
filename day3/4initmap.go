package main

import (
	"fmt"
)

func main() {
	// 1. 字面量初始化
	m1 := map[string]int{"one": 1, "two": 2}
	fmt.Println("1. 字面量:", m1)

	// 2. 使用 make 初始化
	m2 := make(map[string]string)
	m2["name"] = "Alice"
	m2["city"] = "Beijing"
	fmt.Println("2. make 创建:", m2)

	// 3. 声明但未初始化（nil map）
	var m3 map[int]string
	fmt.Println("3. nil map:", m3, "is nil?", m3 == nil)
	// m3[1] = "value" // ❌ panic: assignment to entry in nil map
	m3 = make(map[int]string) // 重新初始化
	m3[1] = "hello"
	fmt.Println("3. 初始化后的 m3:", m3)

	// 4. 使用 new 创建指针 map
	m4 := new(map[string]int)
	*m4 = map[string]int{"x": 100, "y": 200}
	fmt.Println("4. 使用 new:", *m4)

	// 5. 使用 var + 字面量赋值
	var m5 = map[int]bool{1: true, 2: false}
	fmt.Println("5. var + 字面量:", m5)

	// ---------- 对 map 的操作 ----------
	m := map[string]int{"apple": 5, "banana": 3}

	// 增加
	m["cherry"] = 7
	fmt.Println("添加 cherry:", m)

	// 修改
	m["banana"] = 10
	fmt.Println("修改 banana:", m)

	// 查询
	val, ok := m["apple"]
	if ok {
		fmt.Println("apple 存在，值为:", val)
	} else {
		fmt.Println("apple 不存在")
	}

	// 删除
	delete(m, "banana")
	fmt.Println("删除 banana:", m)

	// 遍历
	fmt.Println("遍历 map:")
	for k, v := range m {
		fmt.Printf("%s -> %d\n", k, v)
	}
}
