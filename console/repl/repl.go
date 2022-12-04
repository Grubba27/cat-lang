package repl

import (
	"bufio"
	"cat/console/color"
	"cat/lexer"
	"cat/token"
	"fmt"
	"io"
)

const PROMPT = ">>🐱 "

func Start(i io.Reader, o io.Writer) {
	scanner := bufio.NewScanner(i)
	for {
		_, err :=
			fmt.Fprintf(o, color.Colorize(PROMPT, color.Yellow))

		if err != nil {
			return
		}

		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		for t := l.NextToken(); t.Type != token.EOF; t = l.NextToken() {
			_, err :=
				fmt.Fprintf(o, color.Colorize("%+v\n", color.White), t)

			if err != nil {
				return
			}
		}
	}
}
