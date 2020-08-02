package main

import (
	"fmt"
)

type Message struct {
	Value string
}

func (m Message) sayHi(value string) {
	fmt.Println(m.Value + " " + value)
}

type Greet interface {
	sayHi(value string)
}

func foo(g Greet) {
	g.sayHi("baz")
}

func main() {
	var m Message

	m.Value = "Hello"

	foo(m)
}
