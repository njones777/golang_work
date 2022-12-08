package main
import (
	"fmt"
	"reflect"
)

var p = fmt.Println

func main(){
	//var name type
	//var vName string = "noah"
	//var v1, v2, = 1.2, 3.4 	//asking os to figure out typ
	//var v3 = "hello"
	//v4 := 2.4	//same but you dont have to use the var keyword
	//v4 = 5.4

	p(reflect.TypeOf(25))
	p(reflect.TypeOf(3.24))
	p(reflect.TypeOf(true))
	p(reflect.TypeOf("Hello"))
	p(reflect.TypeOf('A'))



}