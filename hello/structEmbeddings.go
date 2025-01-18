package hello

import "fmt"

type base struct {
	num int
}

func (b base) desribe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

type container struct {
	base
	str string
}

func StructEmbeddings() {
	co := container{
		base: base{num: 1},
		str:  "some string",
	}

	fmt.Println(co.num)
	fmt.Println(co.str)
	fmt.Println(co.base)
	fmt.Println("descript", co.desribe())
}
