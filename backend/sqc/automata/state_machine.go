package automata

import (
	"errors"
	"storemaker-backend/models"
	"sync"
)

// StateMachine - Finite State Machine for workflow management
// Demonstrates FSM for REAL order status transitions from your models

// State represents a state in the FSM (using actual OrderStatus from models)
type State models.OrderStatus

// Event represents a transition trigger
type Event string

// OrderState definitions (mapped from actual models.OrderStatus)
const (
	OrderPending    State = State(models.OrderStatusPending)
	OrderConfirmed  State = State(models.OrderStatusConfirmed)
	OrderProcessing State = State(models.OrderStatusProcessing)
	OrderShipped    State = State(models.OrderStatusShipped)
	OrderDelivered  State = State(models.OrderStatusDelivered)
	OrderCancelled  State = State(models.OrderStatusCancelled)
	OrderRefunded   State = State(models.OrderStatusRefunded)
)

// OrderEvent definitions
const (
	EventConfirm Event = "CONFIRM"
	EventProcess Event = "PROCESS"
	EventShip    Event = "SHIP"
	EventDeliver Event = "DELIVER"
	EventCancel  Event = "CANCEL"
	EventRefund  Event = "REFUND"
)

// Transition represents a state transition
type Transition struct {
	From  State
	Event Event
	To    State
}

// FSM represents a Finite State Machine
type FSM struct {
	current     State
	transitions map[State]map[Event]State
	callbacks   map[Event][]func(from, to State)
	mu          sync.RWMutex
}

// NewFSM creates a new finite state machine
func NewFSM(initialState State) *FSM {
	return &FSM{
		current:     initialState,
		transitions: make(map[State]map[Event]State),
		callbacks:   make(map[Event][]func(from, to State)),
	}
}

// AddTransition adds a valid state transition
func (fsm *FSM) AddTransition(from State, event Event, to State) {
	fsm.mu.Lock()
	defer fsm.mu.Unlock()
	
	if fsm.transitions[from] == nil {
		fsm.transitions[from] = make(map[Event]State)
	}
	fsm.transitions[from][event] = to
}

// Trigger attempts to trigger an event and transition
func (fsm *FSM) Trigger(event Event) error {
	fsm.mu.Lock()
	defer fsm.mu.Unlock()
	
	// Check if transition exists
	if fsm.transitions[fsm.current] == nil {
		return errors.New("no transitions defined for current state")
	}
	
	nextState, ok := fsm.transitions[fsm.current][event]
	if !ok {
		return errors.New("invalid transition")
	}
	
	// Perform transition
	oldState := fsm.current
	fsm.current = nextState
	
	// Execute callbacks
	if callbacks, exists := fsm.callbacks[event]; exists {
		for _, callback := range callbacks {
			callback(oldState, nextState)
		}
	}
	
	return nil
}

// Current returns the current state
func (fsm *FSM) Current() State {
	fsm.mu.RLock()
	defer fsm.mu.RUnlock()
	return fsm.current
}

// CanTransition checks if a transition is valid
func (fsm *FSM) CanTransition(event Event) bool {
	fsm.mu.RLock()
	defer fsm.mu.RUnlock()
	
	if fsm.transitions[fsm.current] == nil {
		return false
	}
	
	_, ok := fsm.transitions[fsm.current][event]
	return ok
}

// OnTransition registers a callback for an event
func (fsm *FSM) OnTransition(event Event, callback func(from, to State)) {
	fsm.mu.Lock()
	defer fsm.mu.Unlock()
	
	if fsm.callbacks[event] == nil {
		fsm.callbacks[event] = make([]func(from, to State), 0)
	}
	fsm.callbacks[event] = append(fsm.callbacks[event], callback)
}

// NewOrderStateMachine creates a pre-configured FSM for orders
// Uses actual order workflow from your storemaker project
func NewOrderStateMachine() *FSM {
	fsm := NewFSM(OrderPending)
	
	// Define valid transitions (based on real order workflow)
	// pending → confirmed → processing → shipped → delivered
	//     ↓          ↓            ↓          ↓
	// cancelled ← ← ← ← ← ← ← ← ← ← ← ← ← ← ← → refunded
	
	fsm.AddTransition(OrderPending, EventConfirm, OrderConfirmed)
	fsm.AddTransition(OrderPending, EventCancel, OrderCancelled)
	
	fsm.AddTransition(OrderConfirmed, EventProcess, OrderProcessing)
	fsm.AddTransition(OrderConfirmed, EventCancel, OrderCancelled)
	
	fsm.AddTransition(OrderProcessing, EventShip, OrderShipped)
	fsm.AddTransition(OrderProcessing, EventCancel, OrderCancelled)
	
	fsm.AddTransition(OrderShipped, EventDeliver, OrderDelivered)
	fsm.AddTransition(OrderShipped, EventCancel, OrderCancelled)
	
	fsm.AddTransition(OrderDelivered, EventRefund, OrderRefunded)
	fsm.AddTransition(OrderCancelled, EventRefund, OrderRefunded)
	
	return fsm
}

// GetValidEvents returns all valid events for the current state
func (fsm *FSM) GetValidEvents() []Event {
	fsm.mu.RLock()
	defer fsm.mu.RUnlock()
	
	events := make([]Event, 0)
	if transitions, ok := fsm.transitions[fsm.current]; ok {
		for event := range transitions {
			events = append(events, event)
		}
	}
	
	return events
}

