package main

import (
	"fmt"

	"github.com/xmlking/hello/stringutil"
	"rsc.io/quote"
)

func init() {
	fmt.Println("Command ==>", "I am Hello")
}

// Hello blablabla
func Hello() string {
	return "Hello, world."
}

func main() {
	fmt.Println(quote.Hello())
	fmt.Println(stringutil.Reverse("!oG ,olleH"))
}
