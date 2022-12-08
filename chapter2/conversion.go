package main

import (
	"fmt"
	"strconv"
)

var pl = fmt.Println
var plf = fmt.Printf

func main() {
	var input string
	fmt.Print("Enter your age:")
	fmt.Scanf("%s", &input)
	age, err := strconv.Atoi(input) //convert string to int
	if err != nil {
		pl(err)
	} else {
		pl("Your age is:", age)
	}
	//Converting string values to specific types
	//Done via the parse functions
	b, err := strconv.ParseBool("t")
	pl("\n", b)    //true
	pl(err)        //nil
	plf("%T\n", b) //bool

	f, err := strconv.ParseFloat("3.1415", 64)
	pl("\n", f)    //3.1415
	pl(err)        //nil
	plf("%T\n", f) //float64

}
