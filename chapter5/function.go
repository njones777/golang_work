package main

import (
	"fmt"
	"time"
)

var pl = fmt.Println

// function with parameters
func displayDate(format string) {
	pl(time.Now().Format(format))
}

func main() {
	displayDate("Mon 2006-01-02 16:09:22")
}
