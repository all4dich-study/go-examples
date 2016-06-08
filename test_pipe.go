package main

import (
	"fmt"
)

func main() {

	c := make(chan int)

	go func(n int) {
		defer close(c)
		for i := 0; i < n; i++ {
			c <- i
		}
	}(10)

	// Pattern: PIPE
	for num := range func(in <-chan int, multi int) <-chan int {
		out := make(chan int)
		go func() {
			defer close(out)
			for num := range in {
				out <- num * multi
			}
		}()
		return out
	}(c, 3) {
		fmt.Println(num)
	}
}
