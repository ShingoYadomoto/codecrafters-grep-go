package main

import (
	"bytes"
	// Uncomment this to pass the first stage
	// "bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode/utf8"
)

// Usage: echo <input_text> | your_grep.sh -E <pattern>
func main() {
	if len(os.Args) < 3 || os.Args[1] != "-E" {
		fmt.Fprintf(os.Stderr, "usage: mygrep -E <pattern>\n")
		os.Exit(2) // 1 means no lines were selected, >1 means error
	}

	pattern := os.Args[2]

	line, err := io.ReadAll(os.Stdin) // assume we're only dealing with a single line
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: read input text: %v\n", err)
		os.Exit(2)
	}

	ok, err := matchLine(line, pattern)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(2)
	}

	if !ok {
		os.Exit(1)
	}

	// default exit code is 0 which means success
}

func matchLine(line []byte, pattern string) (bool, error) {
	if utf8.RuneCountInString(pattern) != 1 && pattern != `\d` && pattern != `\w` && !(strings.HasPrefix(pattern, "[") && strings.HasSuffix(pattern, "]")) {
		return false, fmt.Errorf("unsupported pattern: %q", pattern)
	}

	var ok bool

	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	switch pattern {
	case `\d`:
		chars := "123456789"
		ok = bytes.ContainsAny(line, chars)
	case `\w`:
		alpha := "abcdefghijklmnopqrstuvwxyz"
		chars := "123456789" + alpha + strings.ToUpper(alpha) + "_"
		ok = bytes.ContainsAny(line, chars)
	default:
		ok = bytes.ContainsAny(line, pattern)
	}

	if strings.HasPrefix(pattern, "[") && strings.HasSuffix(pattern, "]") {
		chars := strings.TrimSuffix(strings.TrimPrefix(pattern, "["), "]")
		ok = bytes.ContainsAny(line, chars)
	}

	return ok, nil
}
