package hello

import (
	"fmt"
	"maps"
)

func Maps() {
	m := make(map[string]int)
	m["k1"] = 1
	m["k2"] = 2
	m["k3"] = 3
	m["k4"] = 4
	fmt.Println("map:", m)

	delete(m, "k2")
	fmt.Println("map:", m)
	_, prs := m["k2"] //additional return value when getting a value from a map states if the key exists in the map
	fmt.Println("prs:", prs)
	clear(m)
	fmt.Println("map:", m)

	n := map[string]int{"foo": 1, "bar": 2}
	p := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, p) {
		fmt.Println("m and n are equal")
	} else {
		fmt.Println("m and n are not equal")
	}

}
