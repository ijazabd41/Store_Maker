package sqc

import (
	"fmt"
	"storemaker-backend/sqc/automata"
	"storemaker-backend/sqc/language"
	
	"gorm.io/gorm"
)

// SQC - Software Construction Concepts Module
// Main entry point for the StoreQL query language
// NOW INTEGRATED WITH REAL DATABASE!

// Execute parses and executes a StoreQL query with real database access
func Execute(query string, db *gorm.DB) (*language.QueryResult, error) {
	// Step 1: Lexical Analysis (Automata - DFA)
	lexer := automata.NewLexer(query)
	tokens, err := lexer.Tokenize()
	if err != nil {
		return nil, fmt.Errorf("lexical analysis failed: %w", err)
	}
	
	// Step 2: Parsing (Grammar + ADT Tree)
	parser := language.NewParser(tokens)
	ast, err := parser.Parse()
	if err != nil {
		return nil, fmt.Errorf("parsing failed: %w", err)
	}
	
	// Step 3: Interpretation (Little Language with REAL database)
	interpreter := language.NewInterpreter(ast, db)
	result, err := interpreter.Execute()
	if err != nil {
		return nil, fmt.Errorf("execution failed: %w", err)
	}
	
	return result, nil
}

// Validate checks if a query is syntactically valid
func Validate(query string) error {
	lexer := automata.NewLexer(query)
	tokens, err := lexer.Tokenize()
	if err != nil {
		return err
	}
	
	parser := language.NewParser(tokens)
	_, err = parser.Parse()
	return err
}

// Tokenize returns the tokens for a query (useful for debugging)
func Tokenize(query string) ([]automata.Token, error) {
	lexer := automata.NewLexer(query)
	return lexer.Tokenize()
}

