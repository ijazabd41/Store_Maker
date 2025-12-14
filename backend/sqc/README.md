# Software Construction Concepts Module (SQC)

This module demonstrates five key software construction concepts through a practical **Store Query Language (StoreQL)** implementation.

## Concepts Implemented

### 1. **Abstract Data Types (ADT)**
- **Location**: `adt/`
- **Implementation**: Generic data structures with encapsulation
  - Stack ADT for undo/redo operations
  - Queue ADT for task scheduling
  - Tree ADT for query AST (Abstract Syntax Tree)
  - Set ADT for unique collections

### 2. **Grammar**
- **Location**: `grammar/`
- **Implementation**: Formal grammar definition for StoreQL
  - BNF (Backus-Naur Form) grammar specification
  - Context-free grammar for query parsing
  - Syntax rules for filtering, sorting, and aggregation

### 3. **Automata**
- **Location**: `automata/`
- **Implementation**: Finite State Machines
  - Lexical analyzer (tokenizer) using DFA
  - Order state machine for workflow management
  - Pattern matching automaton for validation

### 4. **Little Language**
- **Location**: `language/`
- **Implementation**: StoreQL - A domain-specific query language
  - Custom syntax for querying stores, products, orders
  - Parser and interpreter
  - Query optimization engine

### 5. **Concurrency**
- **Location**: `concurrency/`
- **Implementation**: Parallel processing patterns
  - Worker pool for batch operations
  - Pipeline pattern for data processing
  - Concurrent query execution
  - Producer-consumer pattern

## StoreQL Query Examples

```sql
-- Find all products in a category with price filter
FIND products WHERE category = "electronics" AND price < 1000 ORDER BY price DESC

-- Get order statistics
ANALYZE orders WHERE status = "completed" GROUP BY date

-- Batch update products
UPDATE products SET discount = 10 WHERE stock > 100 CONCURRENT

-- Complex query with joins
SELECT products, orders 
FROM store "my-store" 
WHERE products.price > 50 AND orders.status = "pending"
LIMIT 10
```

## Architecture

```
sqc/
├── adt/           # Abstract Data Types
├── grammar/       # Grammar definitions
├── automata/      # State machines
├── language/      # StoreQL implementation
├── concurrency/   # Parallel processing
├── examples/      # Usage examples
└── tests/         # Unit tests
```

## API Endpoints

- `POST /api/v1/sqc/query` - Execute StoreQL query
- `POST /api/v1/sqc/validate` - Validate query syntax
- `GET /api/v1/sqc/grammar` - Get grammar specification
- `POST /api/v1/sqc/batch` - Execute batch operations

## Usage

```go
import "storemaker-backend/sqc"

// Execute a query
result, err := sqc.Execute("FIND products WHERE price < 100")

// Parse and validate
ast, err := sqc.Parse("SELECT * FROM products")

// Concurrent batch operation
sqc.BatchUpdate(products, updateFunc, workers: 10)
```

