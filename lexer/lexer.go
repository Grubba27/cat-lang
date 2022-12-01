package lexer

import "cat/token"

// TODO: Support unicode and emojis

type Lexer struct {
	input        string
	position     int  // current position, is the key for the char
	readPosition int  // next position char -> readPosition -> position -> char
	char         byte // char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}

	l.readChar()

	return l
}

func (l *Lexer) readChar() {
	l.char = l.peekChar()
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) skipWhitespace() {
	for l.char == ' ' || l.char == '\t' || l.char == '\n' || l.char == '\r' {
		l.readChar()
	}
}

func newToken(tokenType token.TokenType, char byte) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: string(char),
	}
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func isLetter(char byte) bool {
	// possible variable names are in here [a-z,A-Z,_]
	return 'a' <= char && char <= 'z' || 'A' <= char && char <= 'Z' || char == '_'
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.char) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isDigit(char byte) bool {
	return '0' <= char && char <= '9'
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.char) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.char {

	case '=':
		if l.peekChar() == '=' {
			char := l.char
			l.readChar()
			literal := string(char) + string(l.char)
			tok = token.Token{
				Type:    token.EQ,
				Literal: literal,
			}
		} else {
			tok = newToken(token.ASSIGN, l.char)
		}
	case '+':
		tok = newToken(token.PLUS, l.char)
	case '-':
		tok = newToken(token.MINUS, l.char)
	case '/':
		tok = newToken(token.SLASH, l.char)
	case '*':
		tok = newToken(token.ASTERISK, l.char)

	case '<':
		if l.peekChar() == '=' {
			char := l.char
			l.readChar()
			literal := string(char) + string(l.char)
			tok = token.Token{
				Type:    token.LTE,
				Literal: literal,
			}
		} else {
			tok = newToken(token.LT, l.char)
		}

	case '>':
		if l.peekChar() == '=' {
			char := l.char
			l.readChar()
			literal := string(char) + string(l.char)
			tok = token.Token{
				Type:    token.GTE,
				Literal: literal,
			}
		} else {
			tok = newToken(token.GT, l.char)
		}
	case '(':
		tok = newToken(token.LPAREN, l.char)
	case ')':
		tok = newToken(token.RPAREN, l.char)

	case '{':
		tok = newToken(token.LBRACE, l.char)
	case '}':
		tok = newToken(token.RBRACE, l.char)

	case ',':
		tok = newToken(token.COMMA, l.char)

	case ';':
		tok = newToken(token.SEMICOLON, l.char)

	case '!':
		if l.peekChar() == '=' {
			char := l.char
			l.readChar()
			literal := string(char) + string(l.char)
			tok = token.Token{
				Type:    token.NOT_EQ,
				Literal: literal,
			}
		} else {
			tok = newToken(token.BANG, l.char)
		}
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF

	default:

		if isLetter(l.char) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdentifier(tok.Literal)
			return tok
		}

		if isDigit(l.char) {
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			return tok
		}

		tok = newToken(token.ILLEGAL, l.char)

	}
	l.readChar()
	return tok
}
