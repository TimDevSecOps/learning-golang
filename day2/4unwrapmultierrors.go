package main

import (
	"errors"
	"fmt"
)

// 基础错误
var (
	ErrNetwork = errors.New("network unreachable")
	ErrTimeout = errors.New("request timed out")
)

// 自定义复合错误，表示多个子错误
type MultiError struct {
	Errors []error
}

func (e *MultiError) Error() string {
	return fmt.Sprintf("多个错误: %v", e.Errors)
}

// 实现支持多重错误链
func (e *MultiError) Unwrap() []error {
	return e.Errors
}

func doSomething() error {
	// 模拟两个错误都发生
	return &MultiError{
		Errors: []error{ErrNetwork, ErrTimeout},
	}
}

func main() {
	err := doSomething()

	// 使用 errors.Is 检查其中一个错误是否出现
	if errors.Is(err, ErrTimeout) {
		fmt.Println("检测到超时错误")
	}

	if errors.Is(err, ErrNetwork) {
		fmt.Println("检测到网络错误")
	}

	// 使用 errors.As 捕获结构体类型
	var multiErr *MultiError
	if errors.As(err, &multiErr) {
		fmt.Println("捕获到了 MultiError，包含子错误数量：", len(multiErr.Errors))
	}
}
