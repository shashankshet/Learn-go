package hello

import "fmt"

func Range() {
	nums := []int{2, 3, 4}
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println("total:", total)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Println(k, v)
	}

	for k := range kvs {
		fmt.Println("key:", k)
	}
}
