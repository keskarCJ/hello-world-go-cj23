package main

import (
	"fmt"
	"github.com/hello-world-go-cj23/pkg/util"
	"log"
)

func main() {

	//var message string = "Hello World!";

	if q, e, error := util.CreateMessage(-23); error != nil {
		log.Printf("%v\n", error)
	} else {
		fmt.Printf("%v %v\n", q, e)
	}

	m, v := getMessage()
	a := util.GiveMessage()

	fmt.Printf("%v %v\n", m, v)
	log.Printf("%v\n" ,"[INFO]")
	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", util.HelloWorld)
}

//private
func getMessage() (string, int) {
	return "Hello World!", 1
}

