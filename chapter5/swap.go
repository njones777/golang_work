package main

import "fmt"

var pf = fmt.Printf

//pass by value
func swapv(a, b int) {
	a, b = b, a
}

//swap by address/pointer
func swapa(a, b *int) {
	*a, *b = *b, *a
}

func main() {
	va, vb := 1, 2
	pf("va is %d and vb is %d\n", va, vb)
	swapv(va, vb)
	pf("va is %d and vb is %d\n", va, vb)
	ra, rb := 3, 4
	pf("ra is %d and rb is %d\n", ra, rb)
	swapa(&ra, &rb)
	pf("ra is %d and rb is %d\n", ra, rb)
}
