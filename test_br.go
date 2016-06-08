package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
		wg.Done()
	}()

	wg.Wait()
}
