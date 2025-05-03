package main

import (
	"fmt"
	"time"

	"github.com/TimDevSecOps/learning-golang/day2/p/utils"
)

func main() {
	result := utils.Sum(1, 2)
	fmt.Printf("1 + 2 = %d\n", result)

	now := time.Now()
	fmt.Printf("Current time: %s\n", now.Format("2006-01-02 15:04:05 PM"))
}
