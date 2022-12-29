package parser

import (
	"cat/ast"
	"cat/lexer"
	"testing"
)

func TestLetStatementSimple(t *testing.T) {
	input := `
let x = 5;
	
let y = 10;
	
let foobar = 838383;

`
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("Statements do not contain 3 statements. got=%d",
			len(program.Statements))
	}
	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}
	for index, test := range tests {
		statement := program.Statements[index]
		if !testLetStatement(t, statement, test.expectedIdentifier) {
			return
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {

	if s.TokenLiteral() != "let" {
		t.Errorf("Current token literal is not let. got=%q", s.TokenLiteral())
		return false
	}

	let, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement. got=%T", s)
		return false
	}

	if let.Name.Value != name {
		t.Errorf("letStmt.Name.Value not '%s'. got=%s", name, let.Name.Value)
		return false
	}

	if let.Name.TokenLiteral() != name {
		t.Errorf("letStmt.Name.TokenLiteral() not '%s'. got=%s",
			name, let.Name.TokenLiteral())
		return false
	}

	return true
}
