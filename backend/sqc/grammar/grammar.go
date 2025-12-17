package grammar

// Grammar - Formal grammar definition for StoreQL
// Demonstrates context-free grammar and BNF notation

/*
StoreQL Grammar in BNF (Backus-Naur Form):

<query>         ::= <select-stmt> | <find-stmt> | <update-stmt> | <analyze-stmt>

<select-stmt>   ::= "SELECT" <field-list> "FROM" <source> <where-clause>? <order-clause>? <limit-clause>?

<find-stmt>     ::= "FIND" <entity> <where-clause>? <order-clause>? <limit-clause>?

<update-stmt>   ::= "UPDATE" <entity> "SET" <assignment-list> <where-clause>? <concurrent-clause>?

<analyze-stmt>  ::= "ANALYZE" <entity> <where-clause>? <group-clause>?

<field-list>    ::= <field> | <field> "," <field-list> | "*"

<source>        ::= <entity> | "store" <string>

<entity>        ::= "products" | "orders" | "stores" | "customers"

<where-clause>  ::= "WHERE" <condition>

<condition>     ::= <comparison> | <condition> <logical-op> <condition> | "(" <condition> ")"

<comparison>    ::= <field> <operator> <value>

<operator>      ::= "=" | "!=" | "<" | ">" | "<=" | ">=" | "LIKE" | "IN"

<logical-op>    ::= "AND" | "OR"

<order-clause>  ::= "ORDER BY" <field> <direction>?

<direction>     ::= "ASC" | "DESC"

<limit-clause>  ::= "LIMIT" <number>

<group-clause>  ::= "GROUP BY" <field>

<concurrent-clause> ::= "CONCURRENT" <number>?

<assignment-list> ::= <assignment> | <assignment> "," <assignment-list>

<assignment>    ::= <field> "=" <value>

<field>         ::= <identifier> | <identifier> "." <identifier>

<value>         ::= <string> | <number> | <boolean> | "null"

<string>        ::= '"' <chars> '"' | "'" <chars> "'"

<number>        ::= <digit>+ | <digit>+ "." <digit>+

<boolean>       ::= "true" | "false"

<identifier>    ::= <letter> (<letter> | <digit> | "_")*

<letter>        ::= "a".."z" | "A".."Z"

<digit>         ::= "0".."9"
*/

// GrammarRule represents a production rule in the grammar
type GrammarRule struct {
	NonTerminal string
	Production  []string
}

// Grammar holds the formal grammar definition
type Grammar struct {
	StartSymbol string
	Rules       []GrammarRule
}

// GetStoreQLGrammar returns the complete StoreQL grammar
func GetStoreQLGrammar() *Grammar {
	return &Grammar{
		StartSymbol: "query",
		Rules: []GrammarRule{
			// Query types
			{NonTerminal: "query", Production: []string{"select-stmt"}},
			{NonTerminal: "query", Production: []string{"find-stmt"}},
			{NonTerminal: "query", Production: []string{"update-stmt"}},
			{NonTerminal: "query", Production: []string{"analyze-stmt"}},
			
			// SELECT statement
			{NonTerminal: "select-stmt", Production: []string{"SELECT", "field-list", "FROM", "source"}},
			{NonTerminal: "select-stmt", Production: []string{"SELECT", "field-list", "FROM", "source", "where-clause"}},
			{NonTerminal: "select-stmt", Production: []string{"SELECT", "field-list", "FROM", "source", "where-clause", "order-clause"}},
			
			// FIND statement
			{NonTerminal: "find-stmt", Production: []string{"FIND", "entity"}},
			{NonTerminal: "find-stmt", Production: []string{"FIND", "entity", "where-clause"}},
			{NonTerminal: "find-stmt", Production: []string{"FIND", "entity", "where-clause", "order-clause"}},
			
			// UPDATE statement
			{NonTerminal: "update-stmt", Production: []string{"UPDATE", "entity", "SET", "assignment-list"}},
			{NonTerminal: "update-stmt", Production: []string{"UPDATE", "entity", "SET", "assignment-list", "where-clause"}},
			
			// ANALYZE statement
			{NonTerminal: "analyze-stmt", Production: []string{"ANALYZE", "entity"}},
			{NonTerminal: "analyze-stmt", Production: []string{"ANALYZE", "entity", "where-clause", "group-clause"}},
			
			// Entities
			{NonTerminal: "entity", Production: []string{"products"}},
			{NonTerminal: "entity", Production: []string{"orders"}},
			{NonTerminal: "entity", Production: []string{"stores"}},
			{NonTerminal: "entity", Production: []string{"customers"}},
			
			// Operators
			{NonTerminal: "operator", Production: []string{"="}},
			{NonTerminal: "operator", Production: []string{"!="}},
			{NonTerminal: "operator", Production: []string{"<"}},
			{NonTerminal: "operator", Production: []string{">"}},
			{NonTerminal: "operator", Production: []string{"<="}},
			{NonTerminal: "operator", Production: []string{">="}},
			{NonTerminal: "operator", Production: []string{"LIKE"}},
			{NonTerminal: "operator", Production: []string{"IN"}},
		},
	}
}

// Validate checks if a production rule is valid
func (g *Grammar) Validate() bool {
	// Check if start symbol exists in rules
	hasStartSymbol := false
	for _, rule := range g.Rules {
		if rule.NonTerminal == g.StartSymbol {
			hasStartSymbol = true
			break
		}
	}
	return hasStartSymbol
}

// GetRulesFor returns all production rules for a non-terminal
func (g *Grammar) GetRulesFor(nonTerminal string) []GrammarRule {
	var rules []GrammarRule
	for _, rule := range g.Rules {
		if rule.NonTerminal == nonTerminal {
			rules = append(rules, rule)
		}
	}
	return rules
}

