package hello

import (
	"fmt"
	"math"
)

const pi = 3.14

func Constants() {
	fmt.Println(pi)
	const n = 500000000
	const d = 3e20 / n
	fmt.Println(d)
	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))
}
