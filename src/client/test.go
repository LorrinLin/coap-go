package main

import (
	coap "github.com/lorrin/go-coap"
	"fmt"
)

func main(){
	coap.Dial("udp","localhost:5688")
	fmt.Println("---------------")
}