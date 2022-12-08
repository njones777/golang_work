package main

import "fmt"

var pl = fmt.Println

func main() {
	var num1 = 5 //type infered
	var num2 int //Explicitly typed, since we dont give it a value it will be zero
	fmt.Println(num1)
	pl(num2)

}
