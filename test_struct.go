package main

import (
	"fmt"
	"log"
)

type status string

/* enum : with 'status' */
const (
	UNKNOWN status = "Zero"
	TODO    status = "One"
	DONE    status = "Two"
)

const (
	A = iota
	B
	C
)

type Task struct {
	title  string
	status status
	due    string
}

func main() {
	log.Println("test")
	fmt.Println(UNKNOWN)
	fmt.Println(TODO)
	fmt.Println(DONE)

	var task = Task{
		title:  "Title content",
		status: "check",
		due:    "Due datea",
	}
	fmt.Println(task)
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
}
