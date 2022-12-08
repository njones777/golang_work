////////////////////////////
//Author: Noah Jones
//Program: This is a simple encoder program that I did from a code wars challenge
//		   It takes a sentence or single word and changes the first letter to its ascii value
//		   Along with that it also swaps the second and last character
//Example: Hello
//Outpue:  72olle
////////////////////////////

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//Example run
	str1 := "A wise old owl lived in an oak"
	res := EncryptThis(str1)
	fmt.Println(res)
}
func EncryptThis(text string) string {
	result := ""
	//Check that words were actually provided
	if len(text) > 0 {
		words := strings.Split(text, " ")
		for _, s := range words {
			//If the word is only one letter like A or I
			if len(s) == 1 {
				result += (strconv.Itoa(int(s[0])) + " ")
				//If the word is greater than 2 letters
			} else if len(s) > 2 {
				result += (strconv.Itoa(int(s[0])) + string(s[len(s)-1]) + s[2:len(s)-1] + string(s[1]) + " ")
				//If the word is only two letter like as or an
			} else {
				result += (strconv.Itoa(int(s[0])) + string(s[len(s)-1]) + " ")
			}
		}
		result := strings.TrimRight(result, " ")
		return result

	} else { //If we are given nothing then return nothing
		return ""
	}

}
