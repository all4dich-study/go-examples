package main

import (
	"fmt"
	_ "log"
	_ "os"
	"time"
)

func main() {
	c := make(chan int)
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	go func() {
		for a := 0; a < 5; a++ {
			time.Sleep(1000 * time.Millisecond)
			fmt.Println("in routine")
		}
		c <- 1
	}()
	fmt.Println(list)
	<-c
}
