package main

import (
	"fmt"
	"github.com/uname/siputility"
)

func main() {
	// sip represents recieved SIP message
	sip := []byte{5, 7}
	fmt.Println(siputility.Decode(sip))
}
