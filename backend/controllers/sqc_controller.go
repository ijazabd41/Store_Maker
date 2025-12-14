package controllers

import (
	"net/http"
	"storemaker-backend/sqc"
	"storemaker-backend/sqc/automata"
	"storemaker-backend/sqc/examples"
	"storemaker-backend/sqc/grammar"
	
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SQCController handles Software Construction Concepts endpoints
// NOW INTEGRATED WITH REAL DATABASE!
type SQCController struct{
	db *gorm.DB
}

// NewSQCController creates a new SQC controller with database connection
func NewSQCController(db *gorm.DB) *SQCController {
	return &SQCController{db: db}
}

// ExecuteQuery godoc
// @Summary Execute StoreQL query
// @Description Execute a StoreQL query and return results
// @Tags SQC
// @Accept json
// @Produce json
// @Param request body QueryRequest true "Query to execute"
// @Success 200 {object} QueryResponse
// @Router /api/v1/sqc/query [post]
func (c *SQCController) ExecuteQuery(ctx *gin.Context) {
	var req QueryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	// Execute query with REAL database connection
	result, err := sqc.Execute(req.Query, c.db)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Query execution failed",
			"details": err.Error(),
		})
		return
	}
	
	ctx.JSON(http.StatusOK, QueryResponse{
		Success: result.Success,
		Type:    result.Type,
		Message: result.Message,
		Count:   result.Count,
		Data:    result.Data,
	})
}

// ValidateQuery godoc
// @Summary Validate StoreQL query
// @Description Check if a StoreQL query is syntactically valid
// @Tags SQC
// @Accept json
// @Produce json
// @Param request body QueryRequest true "Query to validate"
// @Success 200 {object} ValidationResponse
// @Router /api/v1/sqc/validate [post]
func (c *SQCController) ValidateQuery(ctx *gin.Context) {
	var req QueryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	err := sqc.Validate(req.Query)
	
	if err != nil {
		ctx.JSON(http.StatusOK, ValidationResponse{
			Valid:   false,
			Message: err.Error(),
		})
		return
	}
	
	ctx.JSON(http.StatusOK, ValidationResponse{
		Valid:   true,
		Message: "Query is syntactically valid",
	})
}

// TokenizeQuery godoc
// @Summary Tokenize StoreQL query
// @Description Get tokens for a StoreQL query (useful for debugging)
// @Tags SQC
// @Accept json
// @Produce json
// @Param request body QueryRequest true "Query to tokenize"
// @Success 200 {object} TokenizeResponse
// @Router /api/v1/sqc/tokenize [post]
func (c *SQCController) TokenizeQuery(ctx *gin.Context) {
	var req QueryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	tokens, err := sqc.Tokenize(req.Query)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Tokenization failed",
			"details": err.Error(),
		})
		return
	}
	
	tokenList := make([]TokenInfo, 0)
	for _, token := range tokens {
		tokenList = append(tokenList, TokenInfo{
			Type:   tokenTypeToString(token.Type),
			Value:  token.Value,
			Line:   token.Line,
			Column: token.Column,
		})
	}
	
	ctx.JSON(http.StatusOK, TokenizeResponse{
		Tokens: tokenList,
		Count:  len(tokenList),
	})
}

// GetGrammar godoc
// @Summary Get StoreQL grammar
// @Description Returns the formal grammar specification for StoreQL
// @Tags SQC
// @Produce json
// @Success 200 {object} GrammarResponse
// @Router /api/v1/sqc/grammar [get]
func (c *SQCController) GetGrammar(ctx *gin.Context) {
	g := grammar.GetStoreQLGrammar()
	
	rules := make([]RuleInfo, 0)
	for _, rule := range g.Rules {
		rules = append(rules, RuleInfo{
			NonTerminal: rule.NonTerminal,
			Production:  rule.Production,
		})
	}
	
	ctx.JSON(http.StatusOK, GrammarResponse{
		StartSymbol: g.StartSymbol,
		Rules:       rules,
		RuleCount:   len(rules),
	})
}

// RunExamples godoc
// @Summary Run SQC examples
// @Description Execute all software construction concept examples with REAL DATA
// @Tags SQC
// @Produce json
// @Success 200 {object} ExamplesResponse
// @Router /api/v1/sqc/examples [get]
func (c *SQCController) RunExamples(ctx *gin.Context) {
	// Run examples with REAL database connection
	examples.RunAllExamples(c.db)
	
	ctx.JSON(http.StatusOK, ExamplesResponse{
		Success: true,
		Message: "Examples executed successfully with REAL DATABASE! Check server console for output.",
		Examples: []string{
			"ADT: Stack, Queue, Tree, Set (Generic data structures)",
			"Automata: Lexer DFA + Order FSM (Using real models.OrderStatus)",
			"Little Language: StoreQL querying REAL products/orders/stores",
			"Grammar: BNF specification for StoreQL",
			"Concurrency: Worker Pool, Pipeline, Producer-Consumer",
		},
	})
}

// GetDocumentation godoc
// @Summary Get SQC documentation
// @Description Returns comprehensive documentation for the SQC module
// @Tags SQC
// @Produce json
// @Success 200 {object} DocumentationResponse
// @Router /api/v1/sqc/docs [get]
func (c *SQCController) GetDocumentation(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, DocumentationResponse{
		Title:       "Software Construction Concepts Module",
		Version:     "1.0.0",
		Description: "Comprehensive implementation of five key software construction concepts",
		Concepts: []ConceptDoc{
			{
				Name:        "Abstract Data Types (ADT)",
				Description: "Generic data structures with encapsulation: Stack, Queue, Tree, Set",
				Location:    "sqc/adt/",
				Examples:    []string{"Stack[T]", "Queue[T]", "Tree[T]", "Set[T]"},
			},
			{
				Name:        "Grammar",
				Description: "Formal BNF grammar definition for StoreQL query language",
				Location:    "sqc/grammar/",
				Examples:    []string{"Context-free grammar", "Production rules", "Syntax validation"},
			},
			{
				Name:        "Automata",
				Description: "Finite State Machines for lexical analysis and workflow management",
				Location:    "sqc/automata/",
				Examples:    []string{"DFA Lexer", "Order State Machine", "Token recognition"},
			},
			{
				Name:        "Little Language",
				Description: "StoreQL - Domain-specific query language with parser and interpreter",
				Location:    "sqc/language/",
				Examples:    []string{"SELECT", "FIND", "UPDATE", "ANALYZE"},
			},
			{
				Name:        "Concurrency",
				Description: "Parallel processing patterns for high-performance operations",
				Location:    "sqc/concurrency/",
				Examples:    []string{"Worker Pool", "Pipeline", "Producer-Consumer", "Fan-In/Fan-Out"},
			},
		},
		QueryExamples: []string{
			"FIND products WHERE price < 100",
			"SELECT * FROM orders WHERE status = \"pending\"",
			"UPDATE products SET discount = 10 WHERE stock > 100",
			"ANALYZE orders WHERE date > \"2024-01-01\" GROUP BY status",
		},
	})
}

// Request/Response types
type QueryRequest struct {
	Query string `json:"query" binding:"required"`
}

type QueryResponse struct {
	Success bool        `json:"success"`
	Type    string      `json:"type"`
	Message string      `json:"message"`
	Count   int         `json:"count"`
	Data    interface{} `json:"data"`
}

type ValidationResponse struct {
	Valid   bool   `json:"valid"`
	Message string `json:"message"`
}

type TokenizeResponse struct {
	Tokens []TokenInfo `json:"tokens"`
	Count  int         `json:"count"`
}

type TokenInfo struct {
	Type   string `json:"type"`
	Value  string `json:"value"`
	Line   int    `json:"line"`
	Column int    `json:"column"`
}

type GrammarResponse struct {
	StartSymbol string     `json:"start_symbol"`
	Rules       []RuleInfo `json:"rules"`
	RuleCount   int        `json:"rule_count"`
}

type RuleInfo struct {
	NonTerminal string   `json:"non_terminal"`
	Production  []string `json:"production"`
}

type ExamplesResponse struct {
	Success  bool     `json:"success"`
	Message  string   `json:"message"`
	Examples []string `json:"examples"`
}

type DocumentationResponse struct {
	Title         string       `json:"title"`
	Version       string       `json:"version"`
	Description   string       `json:"description"`
	Concepts      []ConceptDoc `json:"concepts"`
	QueryExamples []string     `json:"query_examples"`
}

type ConceptDoc struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Location    string   `json:"location"`
	Examples    []string `json:"examples"`
}

// Helper to convert token types
func tokenTypeToString(t automata.TokenType) string {
	types := map[automata.TokenType]string{
		automata.TokenEOF:        "EOF",
		automata.TokenKeyword:    "KEYWORD",
		automata.TokenIdentifier: "IDENTIFIER",
		automata.TokenString:     "STRING",
		automata.TokenNumber:     "NUMBER",
		automata.TokenOperator:   "OPERATOR",
		automata.TokenComma:      "COMMA",
		automata.TokenDot:        "DOT",
		automata.TokenLeftParen:  "LEFT_PAREN",
		automata.TokenRightParen: "RIGHT_PAREN",
		automata.TokenAsterisk:   "ASTERISK",
		automata.TokenError:      "ERROR",
	}
	if name, ok := types[t]; ok {
		return name
	}
	return "UNKNOWN"
}

