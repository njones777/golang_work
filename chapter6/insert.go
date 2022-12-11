// Go doesnt have a built in insert function so you have to implement it yourself
package main

import (
	"errors"
	"fmt"
)

func insert(orig []int, index int, value int) ([]int, error) {
	if index < 0 {
		return nil, errors.New("Index cannot be less than 0")
	}
	if index >= len(orig) {
		return append(orig, value), nil
	}
	orig = append(orig[:index+1], orig[index:]...)
	orig[index] = value
	return orig, nil

}

func main() {
	t := []int{1, 2, 4, 5, 6}
	fmt.Println(t)
	t, err := insert(t, 2, 3)
	fmt.Println(t)
	t, err = insert(t, -1, 7)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t)
}
