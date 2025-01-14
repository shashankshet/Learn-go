package hello

import "fmt"

func Add(a, b int) (int, int) {
	return a + b, a * b
}

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

// closure functions
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
func Function() {
	res, res2 := Add(1, 2)
	fmt.Println(res)
	fmt.Println(res2)
	result := sum(1, 2, 3, 4, 5)
	fmt.Println(result)
	data := []int{1, 2, 3, 4, 5}
	data_res := sum(data...) //syntax to pass slice to variadic function
	fmt.Println(data_res)
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	newInts := intSeq()
	fmt.Println(newInts())
}
