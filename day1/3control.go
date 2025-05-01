package main

import "fmt"

func main() {
	// if-else, if-else if-else
	if score := 85; score >= 90 {
		fmt.Println("A")
	} else if score >= 80 {
		message := "B Grade"
		fmt.Println(message)
	} else {
		fmt.Println("C")
	}
	// fmt.Println("score is:", score): score is not defined
	// fmt.Println("message is:", message): message is not defined

	// for loop
	// 1. standard for:
	for i := 0; i < 5; i++ {
		x := i * 2
		fmt.Println(i, x)
	}
	// fmt.Println(i): i is not defined
	// fmt.Println(x): x is not defined

	// 2. like: while
	count := 0
	for count < 3 {
		fmt.Println("count is:", count)
		count++
	}

	// 3. for range for slice, array, map, string
	nums := []int{1, 2, 3}
	for i, val := range nums {
		fmt.Printf("for range slice: index %d: %d\n", i, val)
	}
	// for range for map
	ages := map[string]int{
		"John": 25,
		"Jane": 30,
	}
	for name, age := range ages {
		fmt.Printf("for range map:name: %s, age: %d\n", name, age)
	}

	// for range for string
	str := "Hello, World!"
	for i, val := range str {
		fmt.Printf("for range string: index %d: %c\n", i, val)
	}

	// switch case
	// fallthrough:
	switch day := 1; day {
	case 1:
		moday := "Monday"
		fmt.Println(moday)
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	default:
		fmt.Println("Unknown day")
	}
	// fmt.Println("day: ", day): day is not defined
	// fmt.Println("moday: ", moday): moday is not defined
}
