package lexer

import "github.com/m-nakamura145/wiig/token"

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

func New(input string) *Lexer {
	I := &Lexer{input: input}
	I.readChar()
	return I
}

func (I *Lexer) readChar() {
	if I.readPosition >= len(I.input) {
		I.ch = 0
	} else {
		I.ch = I.input[I.readPosition]
	}
	I.position = I.readPosition
	I.readPosition += 1
}

func (I *Lexer) NextToken() token.Token {
	var tok token.Token

	I.skipWhitespace()

	switch I.ch {
	case '=':
		if I.peekChar() == '=' {
			ch := I.ch
			I.readChar()
			literal := string(ch) + string(I.ch)
			tok = token.Token{Type: token.EQ, Literal: literal}
		} else {
			tok = newToken(token.ASSIGN, I.ch)
		}
	case ';':
		tok = newToken(token.SEMICOLON, I.ch)
	case '(':
		tok = newToken(token.LPAREN, I.ch)
	case ')':
		tok = newToken(token.RPAREN, I.ch)
	case ',':
		tok = newToken(token.COMMA, I.ch)
	case '+':
		tok = newToken(token.PLUS, I.ch)
	case '-':
		tok = newToken(token.MINUS, I.ch)
	case '!':
		if I.peekChar() == '=' {
			ch := I.ch
			I.readChar()
			literal := string(ch) + string(I.ch)
			tok = token.Token{Type: token.NOT_EQ, Literal: literal}
		} else {
			tok = newToken(token.BANG, I.ch)
		}
	case '/':
		tok = newToken(token.SLASH, I.ch)
	case '*':
		tok = newToken(token.ASTERISK, I.ch)
	case '<':
		tok = newToken(token.LT, I.ch)
	case '>':
		tok = newToken(token.GT, I.ch)
	case '{':
		tok = newToken(token.LBRACE, I.ch)
	case '}':
		tok = newToken(token.RBRACE, I.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(I.ch) {
			tok.Literal = I.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(I.ch) {
			tok.Type = token.INT
			tok.Literal = I.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, I.ch)
		}
	}

	I.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (I *Lexer) skipWhitespace() {
	for I.ch == ' ' || I.ch == '\t' || I.ch == '\n' || I.ch == '\r' {
		I.readChar()
	}
}

func (I *Lexer) readIdentifier() string {
	position := I.position
	for isLetter(I.ch) {
		I.readChar()
	}
	return I.input[position:I.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func (I *Lexer) readNumber() string {
	position := I.position
	for isDigit(I.ch) {
		I.readChar()
	}
	return I.input[position:I.position]
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (I *Lexer) peekChar() byte {
	if I.readPosition >= len(I.input) {
		return 0
	} else {
		return I.input[I.readPosition]
	}
}
