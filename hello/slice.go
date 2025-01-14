package hello

import "fmt"

func Slice() {
	var s []string
	fmt.Println(s, len(s), cap(s))

	s = make([]string, 3)
	fmt.Println(s, len(s), cap(s))
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println(s, len(s), cap(s))
	fmt.Println(s[2])
	s = append(s, "d")
	s = append(s, "e")
	fmt.Println(s)
	c := make([]string, 5)
	copy(c, s)
	fmt.Println(c)
	l := s[2:5]
	fmt.Println(l)
}
