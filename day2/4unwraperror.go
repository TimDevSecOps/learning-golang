package main

import (
	"errors"
	"fmt"
)

// 定义一个底层错误
var ErrDBConn = errors.New("database connection failed")

// 自定义错误类型，用来包装其他错误
type DBError struct {
	Op  string // 操作，比如 "Connect"
	Err error  // 被包装的错误
}

// 实现 error 接口
func (e *DBError) Error() string {
	return fmt.Sprintf("db error during %s: %v", e.Op, e.Err)
}

// 实现 Unwrap，用于 error 链
func (e *DBError) Unwrap() error {
	return e.Err
}

func connectToDB() error {
	// 最底层的错误
	return ErrDBConn
}

func queryData() error {
	err := connectToDB()
	if err != nil {
		// 包装底层错误
		return &DBError{Op: "Connect", Err: err}
	}
	return nil
}

func run() error {
	err := queryData()
	if err != nil {
		// 再次包装错误，形成更高一层错误
		return fmt.Errorf("run failed: %w", err)
	}
	return nil
}

func main() {
	err := run()

	// 判断是否包含最底层错误
	if errors.Is(err, ErrDBConn) {
		fmt.Println("底层数据库连接失败")
	}

	// 提取具体错误类型
	var dbErr *DBError
	if errors.As(err, &dbErr) {
		fmt.Printf("发生数据库错误，操作：%s\n", dbErr.Op)
	}
}
