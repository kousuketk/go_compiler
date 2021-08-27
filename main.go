package main

import (
	"errors"
	"fmt"
	"io/ioutil"
)

type Token struct {
	kind  string // "intliteral", "punct"
	value string
}

var source []byte
var sourceIndex = 0

func getChar() (byte, error) {
	if sourceIndex == len(source) {
		return 0, errors.New("EOF")
	}
	char := source[sourceIndex]
	sourceIndex++
	return char, nil
}

func ungetChar() {
	sourceIndex--
}

func readNumber(char byte) string {
	var number []byte = []byte{char}
	for {
		char, err := getChar()
		if err != nil {
			break
		}
		if '0' <= char && char <= '9' {
			number = append(number, char)
		} else {
			ungetChar()
			break
		}
	}
	return string(number)
}

func tokenize() []*Token {
	var tokens []*Token
	for {
		char, err := getChar()
		if err != nil {
			break
		}
		switch char {
		case ' ', '\n':
			continue
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			number := readNumber(char)
			token := &Token{
				kind:  "intliteral",
				value: number,
			}
			tokens = append(tokens, token)
		case '+', '-', ';':
			token := &Token{
				kind:  "punct",
				value: string([]byte{char}),
			}
			tokens = append(tokens, token)
		default:
			panic(fmt.Sprintf("Invalid cahr '%c'", char))
		}
	}
	return tokens
}

func main() {
	source, _ = ioutil.ReadFile("/dev/stdin")
	tokens := tokenize()
	for _, token := range tokens {
		fmt.Print(token.value, " ")
	}
	fmt.Println()
	// number, err := strconv.Atoi(string(source))
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("\n")
	// fmt.Printf(" .global main\n")
	// fmt.Printf("main:\n")
	// fmt.Printf(" mov $%d, %%rax\n", number)
	// fmt.Printf(" ret\n")
}
