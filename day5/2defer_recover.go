package main

import "fmt"

func main() {
	defer fmt.Println("defer 1")

	defer func() {
		fmt.Println("defer 2 (recovering) start")
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
		fmt.Println("defer 2 (recovering) end")
	}()

	defer fmt.Println("defer 3")

	fmt.Println("About to panic...")
	panic("something went wrong")

	// 这行不会执行
	fmt.Println("This will not be printed")
}
