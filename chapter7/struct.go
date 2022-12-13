package main

import "fmt"

//the fields in a struct don't all have to be the same type
type point struct {
	x float32
	y float32
	z float32
}

func main() {
	//creat of variable of point type
	//var pt1 point
	//each field in the struct can be initialized individually
	/*pt1.x = 3.1
	pt1.y = 5.7
	pt1.z = 4.4*/

	//can also initialize a struct using a struct literal
	pt2 := point{x: 5.6, y: 77.9, z: 8.9}

	fmt.Println(pt2)
	//fmt.Println(pt1.x)
	//fmt.Println(pt1.y)
	//fmt.Println(pt1.z)

}
