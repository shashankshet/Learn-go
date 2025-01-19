package hello

import (
	"fmt"
	"log"
	"net/url"
)

func UrlParsing() {
	s := "postgres://user:pass@host.com:5432/path?k=v#f"
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	fmt.Println(u.Scheme)
	fmt.Println(u.User.Username())

	//logging module
	log.Println("default logger")

	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Println("custom logger")

}
