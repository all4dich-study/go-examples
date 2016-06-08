package main

import (
	"fmt"
	"log"
	"os"
)

type Job struct {
	Command string
	*log.Logger
}

func (job *Job) Logf(format string, args ...interface{}) {
	job.Logger.Printf("%q: %s", job.Command, fmt.Sprintf(format, args...))
}
func main() {
	a := Job{"test", log.New(os.Stderr, "Job:", log.Ldate)}
	a.L

}
