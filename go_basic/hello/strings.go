package hello

import (
	"fmt"
	"unicode/utf8"
)

func String() {
	const s = "สวัสดี"
	fmt.Println(len(s))

	for i := range len(s) {
		fmt.Println(s[i])
	}

	fmt.Println("RuneCountInString:", utf8.RuneCountInString(s))
}
