package main

import (
	"errors"
	"fmt"
)

// 基础错误
var (
	ErrDiskFull         = errors.New("disk full")
	ErrPermissionDenied = errors.New("permission denied")
)

func doFileOperation() error {
	// 模拟多个错误同时发生
	err1 := ErrDiskFull
	err2 := ErrPermissionDenied

	// 使用 errors.Join 合并多个错误
	return errors.Join(err1, err2)
}

func main() {
	err := doFileOperation()

	// 打印合并错误
	fmt.Println("返回的错误:", err)

	// 使用 errors.Is 检查是否包含特定的错误
	if errors.Is(err, ErrDiskFull) {
		fmt.Println("检测到磁盘已满错误")
	}

	if errors.Is(err, ErrPermissionDenied) {
		fmt.Println("检测到权限拒绝错误")
	}

	// 正确方式：断言为 interface{ Unwrap() []error }
	type unwrapper interface {
		Unwrap() []error
	}

	if ue, ok := err.(unwrapper); ok {
		fmt.Println("Unwrap 得到的错误列表:")
		for _, e := range ue.Unwrap() {
			fmt.Printf("- %v\n", e)
		}
	}
}
