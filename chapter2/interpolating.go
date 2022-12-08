package main

import (
	"fmt"
)

var pl = fmt.Println

func main() {
	queue := 5
	name := "Jamie"
	//s := name + ", your queue number is:" + strconv.Itoa(queue)
	//better way
	s := fmt.Sprintf("%s, Your queue number is %d", name, queue)
	pl(s)

}
