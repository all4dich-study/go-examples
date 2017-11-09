package main

// #cgo CFLAGS: -I/Users/sunjoo/work/go/temp
// #cgo LDFLAGS: -L/Users/sunjoo/work/go/temp -lfoo
/*
#include <stdlib.h>
#include "foo.h"
*/
import "C"
import "fmt"

func main() {
	fmt.Println("Hello World")
	fmt.Println(C.PrintMessageSecond())
}
