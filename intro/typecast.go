package main
import (
	"fmt"
)

var p = fmt.Println

func main(){
	cV1 := 1.5
	cV2 := int(cV1)
	p(cV2)
}