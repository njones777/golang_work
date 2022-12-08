package main

import "fmt"

func main() {
	//basic for loop
	/*for i := 10; i > 5; i-- {
		fmt.Println(i)
	}*/
	//In go there is no while loop
	//So you have to use a for loop with just the codition
	x := 99
	for b := 0; b < x; {
		x--
		fmt.Println(x)
	}
}
