package main

/*
#include <stdlib.h>
#include "libtestlib.h"
*/
import "C"
import "fmt"

func main() {
	fmt.Println("Hello World")
	fmt.Println(C.PrintHelloTest())
}
