package examples

import (
	"context"
	"fmt"
	"storemaker-backend/models"
	"storemaker-backend/sqc"
	"storemaker-backend/sqc/adt"
	"storemaker-backend/sqc/automata"
	"storemaker-backend/sqc/concurrency"
	
	"gorm.io/gorm"
)

// Example1_ADT demonstrates Abstract Data Types
func Example1_ADT() {
	fmt.Println("=== ADT Examples ===\n")
	
	// Stack example
	fmt.Println("1. Stack ADT (LIFO):")
	stack := adt.NewStack[string]()
	stack.Push("First")
	stack.Push("Second")
	stack.Push("Third")
	
	fmt.Printf("Stack size: %d\n", stack.Size())
	if val, err := stack.Pop(); err == nil {
		fmt.Printf("Popped: %s\n", val)
	}
	
	// Queue example
	fmt.Println("\n2. Queue ADT (FIFO):")
	queue := adt.NewQueue[int]()
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	
	fmt.Printf("Queue size: %d\n", queue.Size())
	if val, err := queue.Dequeue(); err == nil {
		fmt.Printf("Dequeued: %d\n", val)
	}
	
	// Set example
	fmt.Println("\n3. Set ADT:")
	set := adt.NewSet[string]()
	set.Add("apple")
	set.Add("banana")
	set.Add("apple") // Duplicate, won't be added
	
	fmt.Printf("Set size: %d\n", set.Size())
	fmt.Printf("Contains 'apple': %v\n", set.Contains("apple"))
}

// Example2_Automata demonstrates Finite State Machines with REAL ORDER WORKFLOW
func Example2_Automata() {
	fmt.Println("\n=== Automata Examples (Using Real Order States!) ===\n")
	
	// Lexer example
	fmt.Println("1. Lexer (DFA for tokenization):")
	query := `FIND products WHERE price < 100`
	lexer := automata.NewLexer(query)
	tokens, _ := lexer.Tokenize()
	
	for _, token := range tokens {
		if token.Type != automata.TokenEOF {
			fmt.Printf("Token: %-12s Value: %s\n", tokenTypeToString(token.Type), token.Value)
		}
	}
	
	// State Machine example with REAL order workflow from models
	fmt.Println("\n2. Order State Machine (FSM - REAL ORDER WORKFLOW):")
	fmt.Println("   Using actual OrderStatus from models.Order!")
	fsm := automata.NewOrderStateMachine()
	
	fmt.Printf("Initial state: %s (models.OrderStatusPending)\n", fsm.Current())
	
	fsm.Trigger(automata.EventConfirm)
	fmt.Printf("After CONFIRM: %s (models.OrderStatusConfirmed)\n", fsm.Current())
	
	fsm.Trigger(automata.EventProcess)
	fmt.Printf("After PROCESS: %s (models.OrderStatusProcessing)\n", fsm.Current())
	
	fsm.Trigger(automata.EventShip)
	fmt.Printf("After SHIP: %s (models.OrderStatusShipped)\n", fsm.Current())
	
	fsm.Trigger(automata.EventDeliver)
	fmt.Printf("After DELIVER: %s (models.OrderStatusDelivered)\n", fsm.Current())
}

// Example3_LittleLanguage demonstrates StoreQL with REAL DATABASE
func Example3_LittleLanguage(db *gorm.DB) {
	fmt.Println("\n=== Little Language (StoreQL) Examples - QUERYING REAL DATA! ===\n")
	
	queries := []string{
		"FIND products",   // Will query actual products table
		"FIND orders",     // Will query actual orders table
		"FIND stores",     // Will query actual stores table
		"SELECT * FROM products",
	}
	
	for i, query := range queries {
		fmt.Printf("%d. Query: %s\n", i+1, query)
		
		// Validate syntax
		if err := sqc.Validate(query); err != nil {
			fmt.Printf("   ❌ Invalid: %v\n", err)
		} else {
			fmt.Println("   ✓ Valid syntax")
		}
		
		// Execute on REAL database
		if result, err := sqc.Execute(query, db); err != nil {
			fmt.Printf("   Error: %v\n", err)
		} else {
			fmt.Printf("   Result: %s\n", result.Message)
			
			// Show sample data if available
			switch result.Type {
			case "FIND":
				if result.Data != nil {
					switch data := result.Data.(type) {
					case []models.Product:
						if len(data) > 0 {
							fmt.Printf("   Sample: Product '%s' - $%.2f\n", data[0].Name, data[0].Price)
						}
					case []models.Order:
						if len(data) > 0 {
							fmt.Printf("   Sample: Order #%s - Status: %s\n", data[0].OrderNumber, data[0].Status)
						}
					case []models.Store:
						if len(data) > 0 {
							fmt.Printf("   Sample: Store '%s' - Slug: %s\n", data[0].Name, data[0].Slug)
						}
					}
				}
			}
		}
		fmt.Println()
	}
}

// Example4_Concurrency demonstrates concurrent patterns
func Example4_Concurrency() {
	fmt.Println("\n=== Concurrency Examples ===\n")
	
	// Worker Pool example
	fmt.Println("1. Worker Pool Pattern:")
	pool := concurrency.NewWorkerPool(3, 10)
	pool.Start()
	
	for i := 0; i < 5; i++ {
		taskNum := i
		task := func() error {
			fmt.Printf("   Task %d executed\n", taskNum)
			return nil
		}
		pool.Submit(task)
	}
	
	pool.Stop()
	
	// Pipeline example
	fmt.Println("\n2. Pipeline Pattern:")
	ctx := context.Background()
	
	// Create stages
	filterStage := concurrency.Filter(func(n int) bool { return n%2 == 0 })
	doubleStage := concurrency.Map(func(n int) int { return n * 2 })
	
	// Create pipeline
	numbers := concurrency.Generator(ctx, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	filtered := filterStage(ctx, numbers)
	doubled := doubleStage(ctx, filtered)
	
	fmt.Print("   Even numbers doubled: ")
	for num := range doubled {
		fmt.Printf("%d ", num)
	}
	fmt.Println()
	
	// Producer-Consumer example
	fmt.Println("\n3. Producer-Consumer Pattern:")
	system := concurrency.NewProducerConsumerSystem[int](5)
	
	counter := 0
	produceFunc := func() int {
		counter++
		fmt.Printf("   Produced: %d\n", counter)
		return counter
	}
	
	consumeFunc := func(item int) {
		fmt.Printf("   Consumed: %d\n", item)
	}
	
	// Start producers and consumers
	for i := 0; i < 2; i++ {
		go system.StartProducer(produceFunc)
	}
	
	for i := 0; i < 2; i++ {
		go system.StartConsumer(consumeFunc)
	}
	
	// Let it run briefly
	// In real usage, you'd have proper coordination
	fmt.Println("   (Producer-Consumer running...)")
}

// RunAllExamples runs all demonstration examples with REAL DATABASE
func RunAllExamples(db *gorm.DB) {
	fmt.Println("╔════════════════════════════════════════════════════════╗")
	fmt.Println("║  Software Construction Concepts - INTEGRATED Examples  ║")
	fmt.Println("║  Now using REAL storemaker database!                   ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝\n")
	
	Example1_ADT()
	Example2_Automata()
	Example3_LittleLanguage(db)
	Example4_Concurrency()
	
	fmt.Println("\n╔════════════════════════════════════════════════════════╗")
	fmt.Println("║  All Examples Completed with REAL DATA!                ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝")
}

// Helper function
func tokenTypeToString(t automata.TokenType) string {
	switch t {
	case automata.TokenKeyword:
		return "KEYWORD"
	case automata.TokenIdentifier:
		return "IDENTIFIER"
	case automata.TokenString:
		return "STRING"
	case automata.TokenNumber:
		return "NUMBER"
	case automata.TokenOperator:
		return "OPERATOR"
	default:
		return "OTHER"
	}
}

