package main

import (
	"fmt"
)

var freeList = make(chan *Buffer, 100)
var serverChan = make(chan *Buffer)

func client() {
	for {
		var b *Buffer
		// Grab a buffer if available; allocate if not.
		select {
		case b = <-freeList:
		//Got one; nothing more to do
		default:
			//None free, so allocate a new one.
			b = new(Buffer)
		}
		load(b)
		serverChan <- b
	}
}

func server() {
	for {
		b := <-serverChan
		process(b)
		//Reuse buffer if there's room.
		select {
		case freeList <- b:
		//Buffer on free list; nothing more to do
		default:
			// Free list full, just carry on.
		}
	}
}
func main() {
	fmt.Println("#END:")
}
