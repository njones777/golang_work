package main

import (
	"fmt"
	"time"
)

func displayTime() {
	fmt.Println(time.Now()) //Because this file is in the main package we can call this function in the main.go file
}
