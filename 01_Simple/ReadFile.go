package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	var data, err = ioutil.ReadFile("README.md")

	if err != nil {
		panic(err)
	} else {
		fmt.Println(string(data))
	}
}
