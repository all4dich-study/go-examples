package main

// #cgo CFLAGS: -I/Users/sunjoo/work/go/temp
// #cgo LDFLAGS: -L./temp -lmk_lib
/*
#include <stdlib.h>
#include "foo.h"
*/
import "C"
import "fmt"

func main() {
	fmt.Println("Hello World")
	r := C.PrintMessageSecond()
	fmt.Println(C.GoString(r))
	fmt.Println(*r)
	fmt.Printf("%c\n", *r)
}
