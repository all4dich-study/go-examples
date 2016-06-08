package main

import (
	"fmt"
	"runtime"
)

func main() {
	numCPU := runtime.GOMAXPROCS(0)
	//numCPU := runtime.NumCPU()
	fmt.Println(numCPU)
	for i := 0; i < 3; i++ {
		fmt.Println(i, i*numCPU/3, (i+1)*numCPU/3)
	}
}
