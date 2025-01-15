package hello

import "fmt"

func ZeroVal(ival int) {
	ival = 0
}

func ZeroPtr(iptr *int) {
	*iptr = 0
}

func Pointers() {
	i := 1
	fmt.Println("initial:", i)
	ZeroVal(i)
	fmt.Println("zeroval:", i)
	ZeroPtr(&i)
	fmt.Println("zeroptr:", i)
	fmt.Println("pointer:", &i)
}
