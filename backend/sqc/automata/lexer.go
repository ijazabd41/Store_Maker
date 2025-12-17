package automata

import (
	"strings"
	"unicode"
)

// Lexer - Finite State Automaton for tokenization
// Demonstrates DFA (Deterministic Finite Automaton) implementation

// TokenType represents different token types
type TokenType int

const (
	TokenEOF TokenType = iota
	TokenKeyword
	TokenIdentifier
	TokenString
	TokenNumber
	TokenOperator
	TokenComma
	TokenDot
	TokenLeftParen
	TokenRightParen
	TokenAsterisk
	TokenError
)

// Token represents a lexical token
type Token struct {
	Type    TokenType
	Value   string
	Line    int
	Column  int
}

// Lexer state machine
type Lexer struct {
	input   string
	pos     int
	line    int
	column  int
	state   LexerState
}

// LexerState represents states in the DFA
type LexerState int

const (
	StateStart LexerState = iota
	StateIdentifier
	StateNumber
	StateString
	StateOperator
	StateEnd
)

// Keywords in StoreQL
var keywords = map[string]bool{
	"SELECT": true, "FROM": true, "WHERE": true, "AND": true, "OR": true,
	"ORDER": true, "BY": true, "LIMIT": true, "GROUP": true, "ASC": true, "DESC": true,
	"FIND": true, "UPDATE": true, "SET": true, "ANALYZE": true, "IN": true,
	"LIKE": true, "NULL": true, "TRUE": true, "FALSE": true, "CONCURRENT": true,
}

// NewLexer creates a new lexer instance
func NewLexer(input string) *Lexer {
	return &Lexer{
		input:  input,
		pos:    0,
		line:   1,
		column: 1,
		state:  StateStart,
	}
}

// Tokenize performs lexical analysis and returns all tokens
func (l *Lexer) Tokenize() ([]Token, error) {
	tokens := make([]Token, 0)
	
	for {
		token := l.NextToken()
		tokens = append(tokens, token)
		
		if token.Type == TokenEOF || token.Type == TokenError {
			break
		}
	}
	
	return tokens, nil
}

// NextToken returns the next token using DFA
func (l *Lexer) NextToken() Token {
	// Skip whitespace
	l.skipWhitespace()
	
	if l.pos >= len(l.input) {
		return Token{Type: TokenEOF, Line: l.line, Column: l.column}
	}
	
	ch := l.current()
	startColumn := l.column
	
	// State transitions based on current character
	switch {
	case unicode.IsLetter(ch) || ch == '_':
		return l.scanIdentifierOrKeyword(startColumn)
		
	case unicode.IsDigit(ch):
		return l.scanNumber(startColumn)
		
	case ch == '"' || ch == '\'':
		return l.scanString(ch, startColumn)
		
	case ch == '(':
		l.advance()
		return Token{Type: TokenLeftParen, Value: "(", Line: l.line, Column: startColumn}
		
	case ch == ')':
		l.advance()
		return Token{Type: TokenRightParen, Value: ")", Line: l.line, Column: startColumn}
		
	case ch == ',':
		l.advance()
		return Token{Type: TokenComma, Value: ",", Line: l.line, Column: startColumn}
		
	case ch == '.':
		l.advance()
		return Token{Type: TokenDot, Value: ".", Line: l.line, Column: startColumn}
		
	case ch == '*':
		l.advance()
		return Token{Type: TokenAsterisk, Value: "*", Line: l.line, Column: startColumn}
		
	case isOperatorChar(ch):
		return l.scanOperator(startColumn)
		
	default:
		l.advance()
		return Token{Type: TokenError, Value: string(ch), Line: l.line, Column: startColumn}
	}
}

// DFA state: scanning identifier or keyword
func (l *Lexer) scanIdentifierOrKeyword(startColumn int) Token {
	start := l.pos
	
	// Transition while in identifier state
	for l.pos < len(l.input) && (unicode.IsLetter(l.current()) || unicode.IsDigit(l.current()) || l.current() == '_') {
		l.advance()
	}
	
	value := l.input[start:l.pos]
	valueUpper := strings.ToUpper(value)
	
	// Check if it's a keyword (final state check)
	if keywords[valueUpper] {
		return Token{Type: TokenKeyword, Value: valueUpper, Line: l.line, Column: startColumn}
	}
	
	return Token{Type: TokenIdentifier, Value: value, Line: l.line, Column: startColumn}
}

// DFA state: scanning number
func (l *Lexer) scanNumber(startColumn int) Token {
	start := l.pos
	
	// Integer part
	for l.pos < len(l.input) && unicode.IsDigit(l.current()) {
		l.advance()
	}
	
	// Decimal part
	if l.pos < len(l.input) && l.current() == '.' {
		l.advance()
		for l.pos < len(l.input) && unicode.IsDigit(l.current()) {
			l.advance()
		}
	}
	
	value := l.input[start:l.pos]
	return Token{Type: TokenNumber, Value: value, Line: l.line, Column: startColumn}
}

// DFA state: scanning string
func (l *Lexer) scanString(quote rune, startColumn int) Token {
	l.advance() // Skip opening quote
	start := l.pos
	
	for l.pos < len(l.input) && l.current() != quote {
		l.advance()
	}
	
	if l.pos >= len(l.input) {
		return Token{Type: TokenError, Value: "unterminated string", Line: l.line, Column: startColumn}
	}
	
	value := l.input[start:l.pos]
	l.advance() // Skip closing quote
	
	return Token{Type: TokenString, Value: value, Line: l.line, Column: startColumn}
}

// DFA state: scanning operator
func (l *Lexer) scanOperator(startColumn int) Token {
	start := l.pos
	l.advance()
	
	// Check for two-character operators
	if l.pos < len(l.input) {
		twoChar := l.input[start : l.pos+1]
		if twoChar == "<=" || twoChar == ">=" || twoChar == "!=" || twoChar == "==" {
			l.advance()
			return Token{Type: TokenOperator, Value: twoChar, Line: l.line, Column: startColumn}
		}
	}
	
	value := l.input[start:l.pos]
	return Token{Type: TokenOperator, Value: value, Line: l.line, Column: startColumn}
}

// Helper methods
func (l *Lexer) current() rune {
	if l.pos >= len(l.input) {
		return 0
	}
	return rune(l.input[l.pos])
}

func (l *Lexer) advance() {
	if l.pos < len(l.input) {
		if l.input[l.pos] == '\n' {
			l.line++
			l.column = 1
		} else {
			l.column++
		}
		l.pos++
	}
}

func (l *Lexer) skipWhitespace() {
	for l.pos < len(l.input) && unicode.IsSpace(rune(l.input[l.pos])) {
		l.advance()
	}
}

func isOperatorChar(ch rune) bool {
	return ch == '=' || ch == '!' || ch == '<' || ch == '>' || ch == '+' || ch == '-'
}

