package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

func main() {
	var source []byte
	source, _ = ioutil.ReadFile("/dev/stdin")
	number, err := strconv.Atoi(string(source))
	if err != nil {
		panic(err)
	}
	fmt.Printf("\n")
	fmt.Printf(" .global main\n")
	fmt.Printf("main:\n")
	fmt.Printf(" mov $%d, %%rax\n", number)
	fmt.Printf(" ret\n")
}
