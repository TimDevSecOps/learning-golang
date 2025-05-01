// go mod init github.com/TimDevSecOps/learning-golang
// go mod tidy
// go run day1/1varanduse.go
package main

import "fmt"

const (
	Dev  = iota // 0
	Test        // 1
	Prod        // 2
)

func main() {
	var name string = "ZhangSan" // 显示声明并初始化
	age := 25                    // 类型推导并初始化，只能用在函数内

	var ( // 批量声明
		height float64 = 175.5
		weight int     = 70
	)

	fmt.Printf("name: %s, age: %d, height: %.1f, weight: %d\n", name, age, height, weight)

	// 常量
	const pi = 3.14159
	const (
		red   = 0
		green = 1
		blue  = 2
	)

	// •	iota 是 Go 语言中 用于生成一组自增常量的标识符。
	// •	每遇到一个 const 声明块，iota 会从 0 自动开始递增。
	// •	每行常量声明，如果使用了 iota，就会按行数 +1 的方式自动赋值。
	const (
		sunday    = iota // 0
		monday           // 1
		tuesday          // 2
		wednesday        // 3
		thursday         // 4
		friday           // 5
		saturday         // 6
	)

	const (
		_  = iota             // 0
		KB = 1 << (10 * iota) // 1 << (10 * 1) = 1024
		MB                    // 1 << (10 * 2) = 1024 * 1024
		GB                    // 1 << (10 * 3) = 1024 * 1024 * 1024
		TB                    // 1 << (10 * 4) = 1024 * 1024 * 1024 * 1024
		PB                    // 1 << (10 * 5) = 1024 * 1024 * 1024 * 1024 * 1024
		EB                    // 1 << (10 * 6) = 1024 * 1024 * 1024 * 1024 * 1024 * 1024
	)

	fmt.Printf("KB: %d, MB: %d, GB: %d, TB: %d, PB: %d, EB: %d\n", KB, MB, GB, TB, PB, EB)

	// 状态码枚举（带类型）
	type OrderStatus int

	const (
		Pending   OrderStatus = iota // 0
		Paid                         // 1
		Shipped                      // 2
		Delivered                    // 3
	)

	// 权限掩码（bitmask）
	const (
		CanRead  = 1 << iota // 0001
		CanWrite             // 0010
		CanExec              // 0100
	)
	perm := CanRead | CanWrite // 0011
	fmt.Printf("perm: %d\n", perm)

	const (
		A = iota // 0
		B = 100  // iota 被中断
		C = iota // 2（不是 1！）
	)
	fmt.Printf("A: %d, B: %d, C: %d\n", A, B, C)

	const (
		A1 = iota // 0
		B1 = iota // 1
		C1        // 2
	)
	fmt.Printf("A1: %d, B1: %d, C1: %d\n", A1, B1, C1)

	const (
		A2 = iota // 0
		B2 = 100  // 100
		C2        // 100: 继承上一行的表达式
	)
	fmt.Printf("A2: %d, B2: %d, C2: %d\n", A2, B2, C2)

	const (
		A3 = iota + 1 // 1
		B3            // 2
		C3 = "str"    // "str"
		D3            // "str"
	)
	fmt.Printf("A3: %d, B3: %d, C3: %s, D3: %s\n", A3, B3, C3, D3)

	const (
		A4 = iota    // 0
		B4           // 1
		C4 = "hello" // "hello"
		D4 = iota    // 3 ❗不是 2！因为 iota 始终以“声明行号”为基准
	)
	fmt.Printf("A4: %d, B4: %d, C4: %s, D4: %d\n", A4, B4, C4, D4)
}
