package main

import "fmt"

func main() {
	defer fmt.Println("defer 1")
	defer fmt.Println("defer 2")
	panic("something went wrong")
	// defer 会在 panic 发生后、函数退出前执行
}
