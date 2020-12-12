package main

import (
	"fmt"
	"time"
)

func main() {
	switch time.Now().Weekday() {
	case time.Monday, time.Thursday, time.Wednesday, time.Thursday, time.Friday:
		fmt.Println("It's a weekday")
	default:
		fmt.Println("It's the weekend")
	}

}
