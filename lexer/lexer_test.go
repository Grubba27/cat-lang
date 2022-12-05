package lexer

// TODO: add filenames and linenumbers to tokens
import (
	asserter "cat/test"
	"cat/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
	}

	lexer := New(input)
	for i, test := range tests {
		tok := lexer.NextToken()
		assert := asserter.New(t)

		assert.
			WithIndex(i).
			WithName("tokentype wrong.").
			That(string(tok.Type)).
			ToBe(string(test.expectedType))

		assert.
			WithIndex(i).
			WithName("literal wrong.").
			That(tok.Literal).
			ToBe(test.expectedLiteral)

	}
}
func TestNextTokenSmallProgram(t *testing.T) {

	input := `let five = 5;
let ten = 10;

let add = fn(x,y) {
  x + y;
};

let result = add(five,ten);

`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{

		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},

		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},

		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},

		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},

		{token.LBRACE, "{"},

		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},

		{token.RBRACE, "}"},

		{token.SEMICOLON, ";"},

		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},

		{token.EOF, ""},
	}

	lexer := New(input)
	for i, test := range tests {
		tok := lexer.NextToken()
		assert := asserter.New(t)

		assert.
			WithIndex(i).
			WithName("tokentype wrong.").
			That(string(tok.Type)).
			ToBe(string(test.expectedType))

		assert.
			WithIndex(i).
			WithName("literal wrong.").
			That(tok.Literal).
			ToBe(test.expectedLiteral)
	}
}
func TestNextTokenBooleanOneTokenExpressions(t *testing.T) {

	input := `!-/*5;
5 < 10 > 5;

`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{

		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},

		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.GT, ">"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	lexer := New(input)
	for i, test := range tests {
		tok := lexer.NextToken()
		if tok.Type != test.expectedType {
			t.Fatalf("Tests[%d] - tokentype wrong. expected=%q, got=%q", i, test.expectedType, tok.Type)
		}
		if tok.Literal != test.expectedLiteral {
			t.Fatalf("Tests[%d] - literal wrong. expected=%q, got=%q", i, test.expectedLiteral, tok.Literal)
		}
	}
}
func TestNextTokenBooleanDoubleTokenExpressions(t *testing.T) {

	input := `
10 == 10;
10 != 9;
9 <= 12;
12 >= 9;
`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{

		{token.INT, "10"},
		{token.EQ, "=="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},

		{token.INT, "10"},
		{token.NOT_EQ, "!="},
		{token.INT, "9"},
		{token.SEMICOLON, ";"},

		{token.INT, "9"},
		{token.LTE, "<="},
		{token.INT, "12"},
		{token.SEMICOLON, ";"},

		{token.INT, "12"},
		{token.GTE, ">="},
		{token.INT, "9"},
		{token.SEMICOLON, ";"},

		{token.EOF, ""},
	}

	lexer := New(input)
	for i, test := range tests {
		tok := lexer.NextToken()
		assert := asserter.New(t)

		assert.
			WithIndex(i).
			WithName("tokentype wrong.").
			That(string(tok.Type)).
			ToBe(string(test.expectedType))

		assert.
			WithIndex(i).
			WithName("literal wrong.").
			That(tok.Literal).
			ToBe(test.expectedLiteral)
	}
}
func TestNextTokenIFExpressions(t *testing.T) {

	input := `
if (5 < 10) {
	return true;
} else {
	return false;
}

`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{

		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.EOF, ""},
	}

	lexer := New(input)
	for i, test := range tests {
		tok := lexer.NextToken()
		assert := asserter.New(t)

		assert.
			WithIndex(i).
			WithName("tokentype wrong.").
			That(string(tok.Type)).
			ToBe(string(test.expectedType))

		assert.
			WithIndex(i).
			WithName("literal wrong.").
			That(tok.Literal).
			ToBe(test.expectedLiteral)
	}
}
