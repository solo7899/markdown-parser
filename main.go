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
	BoldToken   = "BOLD"   // **<text>**
	ItalicToken = "ITALIC" // *<text>*
	LinkToken   = "LINK"   // [<name>](<link>)
	ListToken   = "LIST"   // - <text>
	LineToken   = "LINE"
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
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	var lines []string
	// reading line by line
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	//tokeninzing each line
	Lexer(lines)

	//Todo : parsing tokenized file
	newlist := Parser()

	// Todo : writing parsed strings to a file.html
	Write(newlist)

}

// tokenizing each line
func Lexer(lines []string) {
	// defining regular expressions
	bold_expression := re.MustCompile(`\*\*.*\*\*`)
	italic_expression := re.MustCompile(`\*.*\*`)
	link_expression := re.MustCompile(`^\[.*\]\(.*\)$`)
	list_expression := re.MustCompile(`^-\s`)

	// checking lines
	for _, line := range lines {
		if strings.HasPrefix(line, "##") {
			// find Header line
			tokens = append(tokens, Token{Type: HeaderToken, Value: strings.TrimPrefix(line, "##")})
		} else if bold_expression.Match([]byte(line)) {
			// find bold line
			tokens = append(tokens, Token{Type: BoldToken, Value: line})
		} else if italic_expression.Match([]byte(line)) {
			// find italic line
			tokens = append(tokens, Token{Type: ItalicToken, Value: line})
		} else if link_expression.Match([]byte(line)) {
			// find link line
			link_name_start_index := strings.Index(line, "(")
			link_name_end_index := strings.Index(line, ")")
			link_name := line[link_name_start_index+1 : link_name_end_index]

			link_url_start_index := strings.Index(line, "[")
			link_url_end_index := strings.Index(line, "]")
			link_url := line[link_url_start_index+1 : link_url_end_index]

			link := fmt.Sprintf("<a href=\"%s\">%s</a>", link_url, link_name)
			tokens = append(tokens, Token{Type: LineToken, Value: link})
		} else if list_expression.Match([]byte(line)) {
			// find list lines
			tokens = append(tokens, Token{Type: ListToken, Value: strings.TrimPrefix(line, "- ")})
		} else {
			// simple text line
			tokens = append(tokens, Token{Type: LineToken, Value: line})
		}
	}
}

// parsing lines from markdown to html
func Parser() []string {
	var lines []string
	for _, token := range tokens {
		switch token.Type {
		case HeaderToken:
			// translate to Header
			header := fmt.Sprintf("<h1>%s</h1>", token.Value)
			lines = append(lines, header)
		case BoldToken:
			// translate to Bold
			first := strings.Replace(token.Value, "**", "<b>", 1)
			second := strings.Replace(first, "**", "</b>", 1)
			lines = append(lines, second)
		case ItalicToken:
			// to italic
			first := strings.Replace(token.Value, "*", "<i>", 1)
			second := strings.Replace(first, "*", "</i>", 1)
			lines = append(lines, second)
		case LinkToken:
			// to link
			lines = append(lines, token.Value)
		case ListToken:
			// to list
			lines = append(lines, fmt.Sprintf("<ul><li>%s</li></ul>", token.Value))
		default:
			// to simple line
			lines = append(lines, fmt.Sprintf("<p>%s</p>", token.Value))
		}
	}
	return lines
}

func Write(lines []string) {
	// writing to a file

	f, err := os.OpenFile("test.html", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	for _, line := range lines {
		_, err = f.Write([]byte(line))
		if err != nil {
			panic(err)
		}
	}

}
