package main

import (
	"fmt"
	"strconv"
)

// Routine/Channel Examble
func main() {
	noOfRoutines := 4
	numbers := make(chan int)
	done := make(chan bool)
	for i := 0; i < noOfRoutines; i++ {
		go func(j int) {
			for k := range numbers {
				fmt.Println(strconv.Itoa(j) + "/" + strconv.Itoa(k))
			}
			done <- true
		}(i)
	}
	for a := 0; a < 10*100; a++ {
		numbers <- a
	}
	close(numbers)
	for i := 0; i < noOfRoutines; i++ {
		<-done
	}
	fmt.Println("Hello World")
}
