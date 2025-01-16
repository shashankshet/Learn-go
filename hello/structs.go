package hello

import "fmt"

type person struct {
	name    string
	address string
	age     int
}

func NewPerson(name string) *person {
	p := person{name: name}
	p.address = "Earth"
	p.age = 42
	return &p
}

func StructNew() {
	fmt.Println(person{"Bob", "Earth", 42})
	fmt.Println(NewPerson("Alice"))
}
