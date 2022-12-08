package main

import "fmt"

var pl = fmt.Println

func main() {
	firstName := "christopher R. Bartolemew" //can only be used inside functions not outside
	var num1 = 5
	_ = num1 //compiler is now happy even though we dont use this variable
	address := "The address is\n4567 yeye avenue\n 875"
	pl(firstName)
	pl(address)
}
