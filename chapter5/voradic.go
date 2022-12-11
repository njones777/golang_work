//a variabe function takes in a variable number of arguments
//to do this you make a funciton that takes the parameter ...

package main

import "fmt"

func addNums(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

func main() {
	fmt.Println(addNums(1, 4, 5, 6, 7))
	fmt.Println(addNums(10, 7, 8, 5))

}
