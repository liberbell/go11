package main

import (
	"fmt"
	"time"
)

func main() {
	switch {
	case time.Hour() < 12:
		fmt.Println("Before noon")
	default:
		fmt.Println("Afternoon")
	}
}
