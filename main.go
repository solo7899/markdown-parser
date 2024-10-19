package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	HeaderToken = "HEADER"
	BoldToke    = "BOLD"
	ItalicToken = "ITALIC"
	LinkToken   = "LINK"
	ListToken   = "LIST"
)

type Token struct {
	Type  string
	Value string
}

var tokens []Token

func main() {

	// getting file path from arg variable
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Usage : markdown_parser <path/to/your/markdown_file>")
		return
	}

	// opening file
	f, err := os.Open(args[1])
	defer f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	// reading line by line
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		//tokeninzing each line
		Lexer(scanner.Text())
	}

	//Todo : parsing tokenized file
	Parser(tokens)

	// Todo : writing parsed strings to a file.html

}

// tokenizing each line
func Lexer(line string) {

}

// parsing lines from markdown to html
func Parser(tokens []Token) {

}
