package main

import (
	"fmt"
	_ "time"
)

// Sharing a channel
func main() {
	c := make(chan int)
	for i := 0; i < 3; i++ {
		go func(i int) {
			for n := range c {
				//time.Sleep(1)
				fmt.Println(i, n)
			}
		}(i)
	}
	for i := 0; i < 20; i++ {
		c <- i
	}
	close(c)
}
