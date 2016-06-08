package main

import (
	"fmt"
	"log"
	_ "os"
	_ "syscall"
	"time"
)

func main() {
	c := make(chan int)

	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Println("Work failed:", err)
				log.Println(err.)
			}
		}()

		for i := 0; i < 3; i++ {
			fmt.Println(i)
			time.Sleep(1000 * time.Millisecond)
		}
		panic("Paused by a user")
	}()

	fmt.Println("123")
	<-c
}
