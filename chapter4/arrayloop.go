package main

import "fmt"

func main() {
	//array declaration
	var OS [3]string
	OS[0] = "iOS"
	OS[1] = "Windows"
	OS[2] = "Linux"

	//Iterate thorugh array
	for i, v := range OS { //i is the iteration/index number, v is the item at that index
		fmt.Println(i, v)
	}
	//In case you dont care abou the index number
	for _, v := range OS {
		fmt.Println(v)
	}
}
