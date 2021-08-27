package main

import (
	"fmt"
)

func main() {
	fmt.Printf("\n")
	fmt.Printf(" .global main\n")
	fmt.Printf("main:\n")
	fmt.Printf(" mov $42, %%rax\n")
	fmt.Printf(" ret\n")
}
