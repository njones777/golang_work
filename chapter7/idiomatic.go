package main

import "fmt"

type point struct {
	x float32
	y float32
	z float32
}

func newPoint(x, y, z float32) *point {
	p := point{x: x, y: y, z: z}
	return &p
}

func main() {

	p1 := newPoint(68.9, 0.3333, 5.56)
	fmt.Println(p1)
	fmt.Println(p1.x)

}
