package main

import (
	"fmt"
)

func main() {
	c := make(chan struct{}, 100)
	c <- struct{}{}

	close(c)

l:
	for {
		select {
		case _, ok := <-c:
			if !ok {
				break l
			}
			fmt.Println("hja")
		default:
			break l
		}
	}
}
