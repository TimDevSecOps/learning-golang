package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var i int = 43
	var f float64 = 3.14
	var b bool = true
	var s string = "Hello, World!"

	sum := float64(i) + f

	fmt.Printf("int: %d\nfloat: %f\nboolean: %v\nstring: %s\nsum: %f\n", i, f, b, s, sum)

	s = "你好"
	fmt.Println(len(s))                    // 6 bytes
	fmt.Println(utf8.RuneCountInString(s)) // 2: 两个unicode字符

	/*
		fmt format:
		•	%v：通用格式
		•	%#v：Go 语法格式
		•	%T：打印类型
		•	%t：布尔值
		•	%f：浮点型（默认保留6位）
		•	%g：浮点型更紧凑表达
		•	%s：字符串
		•	%q：带双引号的字符串
		•	%x：十六进制表示
		•	%X：十六进制表示，字母大写
		•	%p：指针
		•	%b：二进制表示
	*/

	/*
		zero value:
		•	int: 0
		•	float: 0.000000
		•	bool: false
		•	string: ""
		•	slice: nil
		•	map: nil
		•	pointer: nil
		•	channel: nil
		•	function: nil
		•	interface: nil
		•	map: nil
		•	slice: nil
	*
}
