//Encryption is working right now
//Decryption is still a work in progress

package main

import (
	"fmt"
)

func main() {

	plaintext := "Wearehere"
	x := encrypt(plaintext, 4)
	fmt.Println(x)
	/*ciphertext := encrypt(plaintext)
	fmt.Println(ciphertext)
	decrypt(ciphertext)*/
}

func encrypt(plaintext string, rails int) string {
	//Slice to hold our resulting strings from each different rail/line
	encrypted_text := make([]string, rails)
	counter := 0
	down := false
	for _, value := range plaintext {
		//add value to the encrypted text
		encrypted_text[counter] += string(value)
		//Check if we are at the bottom of the zig-zag and if so change direction
		if counter == rails-1 && down == true {
			counter -= 1
			down = false
			continue
		} else if counter == rails-1 && down == false {
			counter -= 1
			down = true
			continue
		}
		//If we are somewhere in the middle simply itterate to the next position
		if counter != 0 && counter != rails-1 && down == true {
			counter += 1
			continue
		} else if counter != 0 && counter != rails-1 && down == false {
			counter -= 1
			continue
		}
		//Check if we are at the top and if so change direction
		if counter == 0 && down == true {
			counter -= 1
			down = false
			continue
		} else if counter == 0 && down == false {
			counter += 1
			down = true
			continue
		}
	}
	result := ""
	for _, values := range encrypted_text {
		result += values

	}
	return result

} /*
func decrypt(ciphertext string) {
	length := len(ciphertext)
	key := (length / 4)
	fmt.Println("Length is: ", length)
	fmt.Println("Key is: ", key)
	fmt.Println("top is: ", ciphertext[0:key])
	fmt.Println("middle is: ", ciphertext[key:(key*3)])
	fmt.Println("bottom is: ", ciphertext[key*3:length])

}*/
