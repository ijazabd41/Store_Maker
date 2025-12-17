# Software Construction Concepts - Testing Guide

## Quick Start

### Testing via API Endpoints

All SQC features are exposed through REST API endpoints that you can test using any HTTP client.

### Base URL
```
http://localhost:8080/api/v1/sqc
```

## API Endpoints

### 1. Execute Query
**POST** `/sqc/query`

Execute a StoreQL query.

```bash
curl -X POST http://localhost:8080/api/v1/sqc/query \
  -H "Content-Type: application/json" \
  -d '{"query": "FIND products WHERE price < 100"}'
```

Example Response:
```json
{
  "success": true,
  "type": "FIND",
  "message": "Executed FIND products",
  "count": 0,
  "data": null
}
```

### 2. Validate Query
**POST** `/sqc/validate`

Check if a query is syntactically valid.

```bash
curl -X POST http://localhost:8080/api/v1/sqc/validate \
  -H "Content-Type: application/json" \
  -d '{"query": "SELECT * FROM orders WHERE status = \"pending\""}'
```

Example Response:
```json
{
  "valid": true,
  "message": "Query is syntactically valid"
}
```

### 3. Tokenize Query
**POST** `/sqc/tokenize`

Get tokens for debugging.

```bash
curl -X POST http://localhost:8080/api/v1/sqc/tokenize \
  -H "Content-Type: application/json" \
  -d '{"query": "FIND products WHERE price < 100"}'
```

Example Response:
```json
{
  "tokens": [
    {"type": "KEYWORD", "value": "FIND", "line": 1, "column": 1},
    {"type": "IDENTIFIER", "value": "products", "line": 1, "column": 6},
    {"type": "KEYWORD", "value": "WHERE", "line": 1, "column": 15},
    {"type": "IDENTIFIER", "value": "price", "line": 1, "column": 21},
    {"type": "OPERATOR", "value": "<", "line": 1, "column": 27},
    {"type": "NUMBER", "value": "100", "line": 1, "column": 29}
  ],
  "count": 6
}
```

### 4. Get Grammar
**GET** `/sqc/grammar`

Get the formal BNF grammar specification.

```bash
curl http://localhost:8080/api/v1/sqc/grammar
```

### 5. Run Examples
**GET** `/sqc/examples`

Execute all demonstration examples (output in server console).

```bash
curl http://localhost:8080/api/v1/sqc/examples
```

### 6. Get Documentation
**GET** `/sqc/docs`

Get comprehensive module documentation.

```bash
curl http://localhost:8080/api/v1/sqc/docs
```

## Sample Queries

### Valid Queries

1. **FIND Statement**
```sql
FIND products WHERE category = "electronics"
FIND orders WHERE status = "pending"
FIND stores
```

2. **SELECT Statement**
```sql
SELECT * FROM products
SELECT name FROM orders WHERE date > "2024-01-01"
SELECT * FROM stores WHERE active = true
```

3. **UPDATE Statement**
```sql
UPDATE products SET price = 99
UPDATE orders SET status = "shipped"
```

4. **ANALYZE Statement**
```sql
ANALYZE orders WHERE status = "completed"
ANALYZE products
```

## Testing Each Concept

### 1. ADT (Abstract Data Types)
The ADT implementations are used internally by the parser and language components.

Test via:
- Parser creates AST using Tree ADT
- Query execution uses Stack for operator precedence
- Set ADT for unique field collections

### 2. Grammar
Test the formal grammar:
```bash
curl http://localhost:8080/api/v1/sqc/grammar
```

This returns all production rules defined in BNF notation.

### 3. Automata
Test the finite state machine:

**Lexer (DFA)**:
```bash
curl -X POST http://localhost:8080/api/v1/sqc/tokenize \
  -H "Content-Type: application/json" \
  -d '{"query": "SELECT * FROM products WHERE price < 100"}'
```

**State Machine**: The order state machine is demonstrated in the examples endpoint.

### 4. Little Language (StoreQL)
Test the domain-specific language:

```bash
# Parse and execute
curl -X POST http://localhost:8080/api/v1/sqc/query \
  -H "Content-Type: application/json" \
  -d '{"query": "FIND products WHERE price < 100"}'

# Validate syntax
curl -X POST http://localhost:8080/api/v1/sqc/validate \
  -H "Content-Type: application/json" \
  -d '{"query": "SELECT * FROM orders"}'
```

### 5. Concurrency
Concurrency patterns are demonstrated in the examples endpoint:
- Worker Pool
- Pipeline
- Producer-Consumer

```bash
curl http://localhost:8080/api/v1/sqc/examples
```

Check the backend console for output showing concurrent execution.

## PowerShell Test Script

```powershell
# Test all SQC endpoints

$baseUrl = "http://localhost:8080/api/v1/sqc"

Write-Host "=== Testing SQC Module ===" -ForegroundColor Cyan

# Test 1: Execute Query
Write-Host "`n1. Execute Query:" -ForegroundColor Yellow
$response = Invoke-RestMethod -Uri "$baseUrl/query" -Method POST -ContentType "application/json" -Body '{"query": "FIND products WHERE price < 100"}'
$response | ConvertTo-Json

# Test 2: Validate Query
Write-Host "`n2. Validate Query:" -ForegroundColor Yellow
$response = Invoke-RestMethod -Uri "$baseUrl/validate" -Method POST -ContentType "application/json" -Body '{"query": "SELECT * FROM orders"}'
$response | ConvertTo-Json

# Test 3: Tokenize
Write-Host "`n3. Tokenize Query:" -ForegroundColor Yellow
$response = Invoke-RestMethod -Uri "$baseUrl/tokenize" -Method POST -ContentType "application/json" -Body '{"query": "FIND products"}'
$response | ConvertTo-Json -Depth 5

# Test 4: Get Grammar
Write-Host "`n4. Get Grammar:" -ForegroundColor Yellow
$response = Invoke-RestMethod -Uri "$baseUrl/grammar" -Method GET
Write-Host "Rules count: $($response.rule_count)"

# Test 5: Get Documentation
Write-Host "`n5. Get Documentation:" -ForegroundColor Yellow
$response = Invoke-RestMethod -Uri "$baseUrl/docs" -Method GET
Write-Host "Module: $($response.title)"
Write-Host "Version: $($response.version)"
Write-Host "Concepts: $($response.concepts.Count)"

# Test 6: Run Examples
Write-Host "`n6. Run Examples (check server console):" -ForegroundColor Yellow
$response = Invoke-RestMethod -Uri "$baseUrl/examples" -Method GET
$response | ConvertTo-Json

Write-Host "`n=== All Tests Complete ===" -ForegroundColor Green
```

## Expected Outcomes

### Success Indicators
- ✅ Valid queries return `success: true`
- ✅ Invalid queries return validation errors
- ✅ Tokenizer returns correct token list
- ✅ Grammar endpoint returns BNF rules
- ✅ Examples run without errors

### Console Output
When running examples, you should see in the backend console:
- Stack/Queue/Set operations
- Lexer tokenization process
- State machine transitions
- Query parsing and execution
- Concurrent worker activity

## Troubleshooting

### Backend not responding?
1. Check backend is running: `http://localhost:8080/health`
2. Restart backend if needed

### Invalid query errors?
1. Use `/sqc/tokenize` to see how query is parsed
2. Check `/sqc/grammar` for valid syntax
3. Refer to query examples above

### Want to see more?
1. Check server console for detailed output
2. Use `/sqc/examples` to see all concepts in action
3. Review code in `backend/sqc/` directories

## Next Steps

1. **Extend the Language**: Add more keywords and operations
2. **Add Database Integration**: Connect to actual store data
3. **Implement Optimization**: Add query optimization engine
4. **Add More Patterns**: Implement additional concurrency patterns
5. **Create UI**: Build a query editor interface

## Resources

- **Grammar**: `backend/sqc/grammar/grammar.go`
- **ADT Examples**: `backend/sqc/adt/`
- **Automata**: `backend/sqc/automata/`
- **Language**: `backend/sqc/language/`
- **Concurrency**: `backend/sqc/concurrency/`
- **Examples**: `backend/sqc/examples/examples.go`

