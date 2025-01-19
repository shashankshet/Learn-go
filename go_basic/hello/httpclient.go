package hello

import (
	"bufio"
	"fmt"
	"net/http"
)

func HttiCLient() {
	reso, err := http.Get("http://gobyexample.com")
	if err != nil {
		panic(err)
	}
	defer reso.Body.Close()
	fmt.Println("Response status:", reso.Status)

	scanner := bufio.NewScanner(reso.Body)

	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}

}
