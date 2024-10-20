package main

import (
	"bufio"
	"fmt"
	"os"
	re "regexp"
	"strings"
)

// definign Token types
const (
	HeaderToken = "HEADER" // ##<text>
	BoldToke    = "BOLD"   // **<text>**
	ItalicToken = "ITALIC" // *<text>*
	LinkToken   = "LINK"   // [<name>](<link>)
	ListToken   = "LIST"   // - <text>
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

	var lines []string
	// reading line by line
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	//tokeninzing each line
	Lexer(lines)

	//Todo : parsing tokenized file
	Parser(tokens)

	// Todo : writing parsed strings to a file.html

}

// tokenizing each line
func Lexer(lines []string) {
	// defining regular expressions
	bold_expression := re.MustCompile("\\*\\*.*\\*\\*")
	italic_expression := re.MustCompile("\\*.*\\*")
	link_expression := re.MustCompile("\\[.*\\]\\(.*\\)")
	list_expression := re.MustCompile("^-\\s")

	// checking lines
	for _, line := range lines {
		if strings.HasPrefix(line, "##") {
			// find Header line
			tokens = append(tokens, Token{Type: HeaderToken, Value: strings.TrimPrefix(line, "##")})
		} else if bold_expression.Match([]byte(line)) {
			// find bold line

		} else if italic_expression.Match([]byte(line)) {
			// find italic line

		} else if link_expression.Match([]byte(line)) {
			// find link line

		} else if list_expression.Match([]byte(line)) {
			// find list lines

		} else {
			// simple text line

		}
	}
}

// parsing lines from markdown to html
func Parser(tokens []Token) {

}
