package main

import "fmt"

func main() {
	//normal slice
	t := []int{5, 8, 99, 0, 4}
	//to copy you must initialize another slice
	v := make([]int, len(t))
	copy(v, t)
	fmt.Printf("Type of t is %T\n", t)
	fmt.Printf("type of v is %T\n", v)
	fmt.Println(t)
	fmt.Println(v)
}
