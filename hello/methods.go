package hello

type rect struct {
	width, height int
}

func area(r rect) int {
	return r.width * r.height
}

func Methods() {
	r := rect{width: 10, height: 5}
	area := area(r)
	println(area)
}
