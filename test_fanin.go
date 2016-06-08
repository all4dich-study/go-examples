package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sync"
	_ "time"
)

func FanIn(ins ...<-chan string) <-chan string {
	out := make(chan string)
	var wg sync.WaitGroup
	wg.Add(len(ins))
	for _, in := range ins {
		go func(in <-chan string) {
			defer wg.Done()
			for num := range in {
				out <- num
			}
		}(in)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	c3 := make(chan string)
	go func() {
		defer close(c1)
		c1 <- "1"
		c1 <- "2"
	}()
	go func() {
		defer close(c2)
		c2 <- "c"
	}()
	go func() {
		defer close(c3)
		files, err := ioutil.ReadDir(".")
		if err != nil {
			log.Fatal(err)
		}
		for _, file := range files {
			c3 <- file.Name()
		}
	}()
	for v := range FanIn(c1, c2, c3) {
		fmt.Println(v)
	}
	//close(c)
}
