package main

import (
	"fmt"
	smp "github.com/henry0hai/simple-go-unit-test"
)

func main() {
	fmt.Println(callSayHi("Henry Hai"))
}

func callSayHi(name string) string{
	return smp.SayHi(name)
}