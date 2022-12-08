package main

import (
	"fmt"
	"strconv"
)

var pl = fmt.Println

func main() {
	var input1, input2 string
	fmt.Print("provide a number:")
	fmt.Scanf("%s", &input1)
	fmt.Print("Provide another number:")
	fmt.Scanf("%s", &input2)

	//Convert inputs to integers
	Iinput1, err := strconv.Atoi(input1) //convert string to int
	if err != nil {
		pl(err)
	}
	Iinput2, err := strconv.Atoi(input2)
	if err != nil {
		pl(err)
	}

	//Check if numbers are even with one liner
	bool1 := Iinput1%2 == 0
	bool2 := Iinput2%2 == 0

	//Check if both numbers are even
	if bool1 == true && bool2 == true {
		pl("Congrats both numbers are even!!!")
	} else if (bool1 || bool2) && !(bool1 && bool2) { //GO does not have an exclusive or operator, though this is logically equivalent
		pl("One number is odd")
	} else {
		pl("Both Numbers are false")
	}

}
