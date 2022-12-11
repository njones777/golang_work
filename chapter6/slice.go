package main

import "fmt"

func main() {
	//create an empty slice
	s := make([]int, 5) //[0 0 0 0 0]
	fmt.Println(s)
	//create a slice with 2 elements
	s2 := make([]int, 2, 5) //[0 0]
	fmt.Println(s2)
	//Initialize a slice
	t := []int{1, 2, 3, 4, 66}
	fmt.Println(len(t)) //5
	fmt.Println(cap(t)) //5
	//appending to a slice
	t = append(t, 77, 8)
	fmt.Println(t)

}
