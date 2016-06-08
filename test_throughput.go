package main

import (
	"fmt"
)

var MaxOutstanding int = 3
var sem = make(chan int, MaxOutstanding)

func handle(r *Request) {
	sem <- 1
	process(r)
	<-sem
}

func Serve(queue chan *Request) {
	for {
		req := <-queue
		go handle(req)
	}
}

func ServerTwo(queue chan *Request) {
	for req := range queue {
		sem <- 1
		go func(req *Request) {
			process(req)
			<-sem
		}(req)
	}
}

func ServerThree(queue chan *Request) {
	for req := range queue {
		req := req
		sem <- 1
		go func() {
			process(req)
			<-sem
		}()
	}
}

func handle(queue chan *Request) {
	for r := range queue {
		process(r)
	}
}

func ServerFour(clientRequest chan *Request, quit chan bool) {
	for i := 0; i < MaxOutstanding; i++ {
		go handle(clientRequest)
	}
	<-quit
}
