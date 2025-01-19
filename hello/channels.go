package hello

import (
	"fmt"
)

func Channels() {
	message := make(chan string, 2)
	go func() { message <- "ping" }()
	go func() { message <- "pong" }()
	msg := <-message
	fmt.Println(msg)

}
