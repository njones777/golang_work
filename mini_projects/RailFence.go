package main

import "fmt"

func main() {

	plaintext := "Hitherecrazy"

	ciphertext := encrypt(plaintext)
	fmt.Println(ciphertext)
	decrypt(ciphertext)
}

func encrypt(plaintext string) string {
	top, mid, bottom := "", "", ""
	counter := 1
	for _, c := range plaintext {
		//fmt.Println(string(c))
		c := string(c)
		switch counter {
		case 1:
			top += c
			counter++
		case 2, 4:
			mid += c
			counter++
		case 3:
			bottom += c
			counter++
		case 5:
			top += c
			counter = 2
		}
	}
	return (top + mid + bottom)
}
func decrypt(ciphertext string) {
	length := len(ciphertext)
	key := (length / 4)
	fmt.Println("Length is: ", length)
	fmt.Println("Key is: ", key)
	fmt.Println("top is: ", ciphertext[0:key])
	fmt.Println("middle is: ", ciphertext[key:(key*3)])
	fmt.Println("bottom is: ", ciphertext[key*3:length])

}
