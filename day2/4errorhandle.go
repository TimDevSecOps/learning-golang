package main

import (
	"errors"
	"fmt"
)

// 自定义底层数据库错误类型
type DBError struct {
	Query string
	Msg   string
}

// e *DBError: 指针接收者
func (e *DBError) Error() string {
	return fmt.Sprintf("DBError: query=%q msg=%s", e.Query, e.Msg)
}

// Is 方法：判断是否是 DBError 类型
func (e *DBError) Is(target error) bool {
	t, ok := target.(*DBError)
	if !ok {
		return false
	}
	// 判断具体的字段是否相同
	return e.Msg == t.Msg
}

// 模拟一个数据库操作，返回 DBError
func queryDatabase() error {
	return &DBError{
		Query: "SELECT * FROM users WHERE id = 42",
		Msg:   "record not found",
	}
}

// 服务层调用数据库，并封装错误
func getUser() error {
	err := queryDatabase()
	if err != nil {
		// 使用 fmt.Errorf + %w 包装底层错误
		return fmt.Errorf("getUser failed: %w", err)
	}
	return nil
}

// 主函数中判断是否是 DBError 类型
func main() {
	err := getUser()
	if err != nil {
		fmt.Println("Top-level error:", err)

		// 判断是否是 DBError 类型
		var dbErr *DBError
		if errors.As(err, &dbErr) {
			fmt.Println("Root cause is DBError:")
			fmt.Println(" - Query:", dbErr.Query)
			fmt.Println(" - Msg  :", dbErr.Msg)
		} else {
			fmt.Println("Unknown error type.")
		}

		// 判断是否包含具体错误信息（可选）
		targetErr := &DBError{Msg: "record not found"}
		match := errors.Is(err, targetErr)
		fmt.Println("match:", match)
		if match {
			fmt.Println("Matched expected DB error message.")
		}
	}
}
