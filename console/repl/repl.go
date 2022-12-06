package repl

import (
	"bufio"
	"cat/lexer"
	"cat/token"
	"fmt"
	paint "github.com/Grubba27/painter"
	"io"
)

const PROMPT = ">> "

func Start(i io.Reader, o io.Writer) {
	scanner := bufio.NewScanner(i)
	for {
		_, err :=
			fmt.Fprintf(o, paint.InYellow(PROMPT))

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
				fmt.Fprintf(o, paint.InWhite("%+v\n"), t)

			if err != nil {
				return
			}
		}
	}
}
