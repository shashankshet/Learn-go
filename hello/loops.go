package hello

import "fmt"

func Loops() {
	i := 1
	//single condition loop
	for i < 3 {
		fmt.Println(i)
		i = i + 1
	}

	//classic initial/condition/after loop
	for j := 0; j < 10; j++ {
		fmt.Println(j)
	}

	//using range to iterate over loop
	for i := range 10 {
		fmt.Println(i)
	}

	//infinite loop
	for {
		fmt.Println("loop")
		break
	}
}
