package main
import (
	"fmt"
	"bufio"
	"os"
	"log"
)
var pl = fmt.Println
func main(){
	pl("What is your name")
	//sets up a buffer reader to get text from keyboard
	reader := bufio.NewReader(os.Stdin)	
	name, err := reader.ReadString('\n') //read up until new line
	
	//if nil is passed back no error was thrown
	if err == nil { 
		pl("Hello ", name)
	} else{
		log.Fatal(err)
	}
}