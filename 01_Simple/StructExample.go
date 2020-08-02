package main

import (
	"fmt"
)

type Message struct {
	Value string
}

// Bound function to struct
func (m Message) sayHi(value string) {
	fmt.Println(m.Value + " " + value)
}

func main() {

	var m Message

	m.Value = "Hello"

	m.sayHi(" World!!!")
}
