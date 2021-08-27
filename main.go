package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
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

var tokens []*Token
var tokenIndex = 0

func getToken() *Token {
	if tokenIndex == len(tokens) {
		return nil
	}
	token := tokens[tokenIndex]
	tokenIndex++
	return token
}

type Expr struct {
	kind     string // intliteral", "unary"
	intval   int
	operator string // "+", "-"
	operand  *Expr  // 式の中に式がある、再帰的な構造になっている
}

func parse() *Expr {
	token := getToken()
	switch token.kind {
	case "intliteral":
		number, err := strconv.Atoi(token.value)
		if err != nil {
			panic(err)
		}
		return &Expr{
			kind:   "intliteral",
			intval: number,
		}
	case "punct":
		return &Expr{
			kind:     "unary",
			operator: token.value,
			operand:  parse(),
		}
	default:
		panic("Unexpected token.kind")
	}
}

func generateExpr(expr *Expr) {
	switch expr.kind {
	case "intliteral":
		fmt.Printf(" mov $%d, %%rax\n", expr.intval)
	case "unary":
		switch expr.operator {
		case "+":
			fmt.Printf(" mov $%d, %%rax\n", expr.operand.intval)
		case "-":
			fmt.Printf(" mov $-%d, %%rax\n", expr.operand.intval)
		}
	default:
		panic("Unexpected expr.kind")
	}
}

func main() {
	source, _ = ioutil.ReadFile("/dev/stdin")
	tokens = tokenize()
	expr := parse()

	fmt.Printf("\n")
	fmt.Printf(" .global main\n")
	fmt.Printf("main:\n")
	generateExpr(expr)
	fmt.Printf(" ret\n")
}
