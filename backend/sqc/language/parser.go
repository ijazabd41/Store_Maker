package language

import (
	"errors"
	"fmt"
	"storemaker-backend/sqc/automata"
	"storemaker-backend/sqc/adt"
)

// Parser - Converts tokens to Abstract Syntax Tree
// Implements the Little Language (StoreQL) parser

// ASTNode types
type NodeType string

const (
	NodeQuery      NodeType = "QUERY"
	NodeSelect     NodeType = "SELECT"
	NodeFind       NodeType = "FIND"
	NodeUpdate     NodeType = "UPDATE"
	NodeAnalyze    NodeType = "ANALYZE"
	NodeWhere      NodeType = "WHERE"
	NodeCondition  NodeType = "CONDITION"
	NodeField      NodeType = "FIELD"
	NodeValue      NodeType = "VALUE"
	NodeOperator   NodeType = "OPERATOR"
	NodeOrderBy    NodeType = "ORDER_BY"
	NodeLimit      NodeType = "LIMIT"
)

// ASTNode represents a node in the Abstract Syntax Tree
type ASTNode struct {
	Type     NodeType
	Value    string
	Children []*ASTNode
	Metadata map[string]interface{}
}

// Parser parses tokens into AST
type Parser struct {
	tokens []automata.Token
	pos    int
	ast    *adt.TreeNode[*ASTNode]
}

// NewParser creates a new parser
func NewParser(tokens []automata.Token) *Parser {
	return &Parser{
		tokens: tokens,
		pos:    0,
	}
}

// Parse converts tokens to AST
func (p *Parser) Parse() (*adt.TreeNode[*ASTNode], error) {
	if len(p.tokens) == 0 {
		return nil, errors.New("no tokens to parse")
	}
	
	// Create root node
	root := &adt.TreeNode[*ASTNode]{
		Value:    &ASTNode{Type: NodeQuery, Value: "ROOT"},
		Children: make([]*adt.TreeNode[*ASTNode], 0),
	}
	p.ast = root
	
	// Parse query based on first keyword
	if err := p.parseQuery(); err != nil {
		return nil, err
	}
	
	return root, nil
}

// parseQuery identifies and parses the query type
func (p *Parser) parseQuery() error {
	if !p.hasMore() {
		return errors.New("unexpected end of input")
	}
	
	token := p.currentToken()
	
	switch token.Value {
	case "SELECT":
		return p.parseSelectStatement()
	case "FIND":
		return p.parseFindStatement()
	case "UPDATE":
		return p.parseUpdateStatement()
	case "ANALYZE":
		return p.parseAnalyzeStatement()
	default:
		return fmt.Errorf("unexpected keyword: %s", token.Value)
	}
}

// parseSelectStatement parses SELECT queries
func (p *Parser) parseSelectStatement() error {
	selectNode := p.addNode(NodeSelect, "SELECT")
	p.advance() // Skip SELECT
	
	// Parse field list
	if err := p.parseFieldList(selectNode); err != nil {
		return err
	}
	
	// Expect FROM
	if !p.expect("FROM") {
		return errors.New("expected FROM keyword")
	}
	p.advance()
	
	// Parse source
	if err := p.parseSource(selectNode); err != nil {
		return err
	}
	
	// Optional WHERE clause
	if p.peek("WHERE") {
		if err := p.parseWhereClause(selectNode); err != nil {
			return err
		}
	}
	
	// Optional ORDER BY clause
	if p.peek("ORDER") {
		if err := p.parseOrderByClause(selectNode); err != nil {
			return err
		}
	}
	
	// Optional LIMIT clause
	if p.peek("LIMIT") {
		if err := p.parseLimitClause(selectNode); err != nil {
			return err
		}
	}
	
	return nil
}

// parseFindStatement parses FIND queries
func (p *Parser) parseFindStatement() error {
	findNode := p.addNode(NodeFind, "FIND")
	p.advance() // Skip FIND
	
	// Parse entity
	if p.currentToken().Type != automata.TokenIdentifier && p.currentToken().Type != automata.TokenKeyword {
		return errors.New("expected entity name")
	}
	
	entityNode := &adt.TreeNode[*ASTNode]{
		Value:    &ASTNode{Type: NodeField, Value: p.currentToken().Value},
		Children: make([]*adt.TreeNode[*ASTNode], 0),
		Parent:   findNode,
	}
	findNode.Children = append(findNode.Children, entityNode)
	p.advance()
	
	// Optional WHERE clause
	if p.peek("WHERE") {
		if err := p.parseWhereClause(findNode); err != nil {
			return err
		}
	}
	
	// Optional ORDER BY
	if p.peek("ORDER") {
		if err := p.parseOrderByClause(findNode); err != nil {
			return err
		}
	}
	
	return nil
}

// parseUpdateStatement parses UPDATE queries
func (p *Parser) parseUpdateStatement() error {
	updateNode := p.addNode(NodeUpdate, "UPDATE")
	p.advance() // Skip UPDATE
	
	// Parse entity
	if p.currentToken().Type != automata.TokenIdentifier {
		return errors.New("expected entity name")
	}
	
	entityNode := &adt.TreeNode[*ASTNode]{
		Value:    &ASTNode{Type: NodeField, Value: p.currentToken().Value},
		Children: make([]*adt.TreeNode[*ASTNode], 0),
		Parent:   updateNode,
	}
	updateNode.Children = append(updateNode.Children, entityNode)
	p.advance()
	
	// Expect SET
	if !p.expect("SET") {
		return errors.New("expected SET keyword")
	}
	p.advance()
	
	// Parse assignments
	// This is simplified - you would parse field = value pairs
	
	return nil
}

// parseAnalyzeStatement parses ANALYZE queries
func (p *Parser) parseAnalyzeStatement() error {
	analyzeNode := p.addNode(NodeAnalyze, "ANALYZE")
	p.advance() // Skip ANALYZE
	
	// Parse entity
	if p.currentToken().Type != automata.TokenIdentifier {
		return errors.New("expected entity name")
	}
	
	entityNode := &adt.TreeNode[*ASTNode]{
		Value:    &ASTNode{Type: NodeField, Value: p.currentToken().Value},
		Children: make([]*adt.TreeNode[*ASTNode], 0),
		Parent:   analyzeNode,
	}
	analyzeNode.Children = append(analyzeNode.Children, entityNode)
	p.advance()
	
	return nil
}

// parseFieldList parses field names
func (p *Parser) parseFieldList(parent *adt.TreeNode[*ASTNode]) error {
	// Simplified: just accept identifier or *
	token := p.currentToken()
	
	if token.Type == automata.TokenAsterisk {
		fieldNode := &adt.TreeNode[*ASTNode]{
			Value:    &ASTNode{Type: NodeField, Value: "*"},
			Children: make([]*adt.TreeNode[*ASTNode], 0),
			Parent:   parent,
		}
		parent.Children = append(parent.Children, fieldNode)
		p.advance()
	} else if token.Type == automata.TokenIdentifier {
		fieldNode := &adt.TreeNode[*ASTNode]{
			Value:    &ASTNode{Type: NodeField, Value: token.Value},
			Children: make([]*adt.TreeNode[*ASTNode], 0),
			Parent:   parent,
		}
		parent.Children = append(parent.Children, fieldNode)
		p.advance()
	}
	
	return nil
}

// parseSource parses FROM source
func (p *Parser) parseSource(parent *adt.TreeNode[*ASTNode]) error {
	if p.currentToken().Type != automata.TokenIdentifier && p.currentToken().Type != automata.TokenKeyword {
		return errors.New("expected source name")
	}
	
	sourceNode := &adt.TreeNode[*ASTNode]{
		Value:    &ASTNode{Type: NodeField, Value: p.currentToken().Value},
		Children: make([]*adt.TreeNode[*ASTNode], 0),
		Parent:   parent,
	}
	parent.Children = append(parent.Children, sourceNode)
	p.advance()
	
	return nil
}

// parseWhereClause parses WHERE conditions
func (p *Parser) parseWhereClause(parent *adt.TreeNode[*ASTNode]) error {
	p.advance() // Skip WHERE
	
	whereNode := &adt.TreeNode[*ASTNode]{
		Value:    &ASTNode{Type: NodeWhere, Value: "WHERE"},
		Children: make([]*adt.TreeNode[*ASTNode], 0),
		Parent:   parent,
	}
	parent.Children = append(parent.Children, whereNode)
	
	// Simplified: field operator value
	// In full implementation, handle complex boolean expressions
	
	return nil
}

// parseOrderByClause parses ORDER BY
func (p *Parser) parseOrderByClause(parent *adt.TreeNode[*ASTNode]) error {
	p.advance() // Skip ORDER
	
	if !p.expect("BY") {
		return errors.New("expected BY after ORDER")
	}
	p.advance()
	
	orderNode := &adt.TreeNode[*ASTNode]{
		Value:    &ASTNode{Type: NodeOrderBy, Value: "ORDER_BY"},
		Children: make([]*adt.TreeNode[*ASTNode], 0),
		Parent:   parent,
	}
	parent.Children = append(parent.Children, orderNode)
	
	return nil
}

// parseLimitClause parses LIMIT
func (p *Parser) parseLimitClause(parent *adt.TreeNode[*ASTNode]) error {
	p.advance() // Skip LIMIT
	
	if p.currentToken().Type != automata.TokenNumber {
		return errors.New("expected number after LIMIT")
	}
	
	limitNode := &adt.TreeNode[*ASTNode]{
		Value:    &ASTNode{Type: NodeLimit, Value: p.currentToken().Value},
		Children: make([]*adt.TreeNode[*ASTNode], 0),
		Parent:   parent,
	}
	parent.Children = append(parent.Children, limitNode)
	p.advance()
	
	return nil
}

// Helper methods
func (p *Parser) hasMore() bool {
	return p.pos < len(p.tokens) && p.tokens[p.pos].Type != automata.TokenEOF
}

func (p *Parser) currentToken() automata.Token {
	if p.pos < len(p.tokens) {
		return p.tokens[p.pos]
	}
	return automata.Token{Type: automata.TokenEOF}
}

func (p *Parser) advance() {
	if p.pos < len(p.tokens) {
		p.pos++
	}
}

func (p *Parser) peek(value string) bool {
	return p.hasMore() && p.currentToken().Value == value
}

func (p *Parser) expect(value string) bool {
	return p.hasMore() && p.currentToken().Value == value
}

func (p *Parser) addNode(nodeType NodeType, value string) *adt.TreeNode[*ASTNode] {
	node := &adt.TreeNode[*ASTNode]{
		Value:    &ASTNode{Type: nodeType, Value: value},
		Children: make([]*adt.TreeNode[*ASTNode], 0),
		Parent:   p.ast,
	}
	p.ast.Children = append(p.ast.Children, node)
	return node
}

