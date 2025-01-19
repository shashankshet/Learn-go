package hello

import "fmt"

func f(arg int) (int, error) {
	if arg == 42 {
		return -1, fmt.Errorf("can't work with 42")
	}
	return arg + 3, nil
}

var erorNoTea = fmt.Errorf("no tea")
var errorPower = fmt.Errorf("no power")

func makeTea(arg int) (string, error) {
	if arg == 0 {
		return "", erorNoTea
	} else if arg == 1 {
		return "", errorPower
	}
	return "tea", nil
}

func Errors() {
	fmt.Println(f(42))
	fmt.Println(makeTea(0))
}
