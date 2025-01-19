package hello

import (
	"fmt"
	"time"
)

func Switch() {
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		{
			fmt.Println("It's the weekend")
		}
	default:
		{
			fmt.Println("It's a weekday")
		}
	}
}
