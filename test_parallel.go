package main

import (
	"fmt"
)

type Vector []float64

const numCPU = 4

// Apply the operation to v[i], v[i+1] ... up to v[n-1]
func (v Vector) DoSome(i, n int, u Vector, c chan int) {
	for ; i < n; i++ {
		v[i] += u.Op(v[i])
	}
	c <- 1 //Insert a data into a channel
}

func (v Vector) DoAll(u Vector) {
	c := make(chan int, numCPU)
	for i := 0; i < numCPU; i++ {
		go v.DoSome(i*len(v)/numCPU, (i+1)*len(v)/numCPU, u, c)
	}

	for i := 0; i < numCPU; i++ {
		<-c //Extract a data from a channel
	}
}

func main() {
	fmt.Println("123")
}
