package main

import "fmt"

func divide(a, b int) int {
	return a / b
}

func riskyOperation() {
	fmt.Println("About to divide by zero inside riskyOperation start")
	_ = divide(10, 0)
	fmt.Println("About to divide by zero inside riskyOperation end")
}

// output:
// Starting risky operation
// About to divide by zero inside riskyOperation start
// defer 2 start
// defer 2 recovered: runtime error: integer divide by zero
// defer 1 start
// defer 1 recovered: panic from defer 2
// defer 1 end
func main() {
	defer func() {
		fmt.Println("defer 1 start")
		if r := recover(); r != nil {
			fmt.Println("defer 1 recovered:", r)
		}
		fmt.Println("defer 1 end")
	}()

	defer func() {
		fmt.Println("defer 2 start")
		if r := recover(); r != nil {
			fmt.Println("defer 2 recovered:", r)
		}
		panic("panic from defer 2")
		fmt.Println("defer 2 end (unreachable)")
	}()

	fmt.Println("Starting risky operation")
	riskyOperation()

	panic("original panic")
}
