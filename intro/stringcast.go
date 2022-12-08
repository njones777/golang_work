package main
import (
	"fmt"
	"strconv"
	"reflect"
)
var pl = fmt.Println
func main(){
	//String to integer
	cV3 := "50000000"
	cV4, err := strconv.Atoi(cV3) //Atoi means Ascii to integer
	pl(cV4, err, reflect.TypeOf(cV4))

	//Integer to string
	cV6 := strconv.Itoa(cV4)	//Itoa means Integer to ASCII
	pl(cV4, err, reflect.TypeOf(cV6))

	//string to float
	cV7 := "3.14"
	if cV8, err := strconv.ParseFloat(cV7, 64); err == nil{
		pl(cV8)
	}
	//formated printing
	cV9 := fmt.Sprintf("%f",2.1189)
	pl(reflect.TypeOf(cV9))
}