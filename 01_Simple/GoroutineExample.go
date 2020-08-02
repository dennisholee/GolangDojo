package main

// TODO: Fix early completion bug

import (
	"fmt"
	"sync"
)

func PrintRepeat(wg *sync.WaitGroup, c chan string, v string) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		c <- v
	}
}

func main() {
	var wg sync.WaitGroup
	var c = make(chan string, 1)

	wg.Add(3)
	go PrintRepeat(&wg, c, "a")
	go PrintRepeat(&wg, c, "b")
	go PrintRepeat(&wg, c, "c")

	go func() {
		for data := range c {
			fmt.Print(data)
		}
	}()

	wg.Wait()
}
