package parser

import (
	"cat/ast"
	"cat/lexer"
	"cat/token"
)

type Parser struct {
	lexer        *lexer.Lexer
	currentToken token.Token
	nextToken    token.Token
}

func New(lexer *lexer.Lexer) *Parser {
	p := &Parser{lexer: lexer}
	p.peekNext()
	p.peekNext()
	return p

}
func (p *Parser) currentTokenIs(t token.TokenType) bool {
	return p.currentToken.Type == t
}
func (p *Parser) nextTokenIs(t token.TokenType) bool {
	return p.nextToken.Type == t
}
func (p *Parser) expectNextToBe(t token.TokenType) bool {
	if p.nextTokenIs(t) {
		p.peekNext()
		return true
	}
	return false
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	statement := &ast.LetStatement{Token: p.currentToken}
	if !p.expectNextToBe(token.IDENT) {
		return nil
	}
	statement.Name = &ast.Identifier{Token: p.currentToken, Value: p.currentToken.Literal}
	if !p.expectNextToBe(token.ASSIGN) {
		return nil
	}

	for !p.currentTokenIs(token.SEMICOLON) {
		p.peekNext()
	}
	return statement
}
func (p *Parser) peekNext() {
	p.currentToken = p.nextToken
	p.nextToken = p.lexer.NextToken()
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.currentToken.Type {
	case token.LET:
		return p.parseLetStatement()
	default:
		return nil
	}
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}
	for p.currentToken.Type != token.EOF {
		statement := p.parseStatement()

		if statement != nil {
			program.Statements = append(program.Statements, statement)
		}
		p.peekNext()
	}
	return program
}
