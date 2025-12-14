package adt

// Tree ADT - Used for representing Abstract Syntax Trees (AST)
// Essential for parsing and representing query structure

// TreeNode represents a node in the AST
type TreeNode[T any] struct {
	Value    T
	Children []*TreeNode[T]
	Parent   *TreeNode[T]
}

// Tree interface defines tree operations
type Tree[T any] interface {
	Root() *TreeNode[T]
	Insert(parent *TreeNode[T], value T) *TreeNode[T]
	Traverse(visitor func(node *TreeNode[T]))
	Height() int
	Size() int
}

// GenericTree implements a general tree structure
type GenericTree[T any] struct {
	root *TreeNode[T]
	size int
}

// NewTree creates a new tree with a root value
func NewTree[T any](rootValue T) Tree[T] {
	return &GenericTree[T]{
		root: &TreeNode[T]{Value: rootValue, Children: make([]*TreeNode[T], 0)},
		size: 1,
	}
}

// Root returns the root node
func (t *GenericTree[T]) Root() *TreeNode[T] {
	return t.root
}

// Insert adds a child node to the specified parent
func (t *GenericTree[T]) Insert(parent *TreeNode[T], value T) *TreeNode[T] {
	newNode := &TreeNode[T]{
		Value:    value,
		Children: make([]*TreeNode[T], 0),
		Parent:   parent,
	}
	
	if parent != nil {
		parent.Children = append(parent.Children, newNode)
	}
	
	t.size++
	return newNode
}

// Traverse performs depth-first traversal
func (t *GenericTree[T]) Traverse(visitor func(node *TreeNode[T])) {
	if t.root != nil {
		t.traverseHelper(t.root, visitor)
	}
}

func (t *GenericTree[T]) traverseHelper(node *TreeNode[T], visitor func(node *TreeNode[T])) {
	visitor(node)
	for _, child := range node.Children {
		t.traverseHelper(child, visitor)
	}
}

// Height returns the height of the tree
func (t *GenericTree[T]) Height() int {
	return t.heightHelper(t.root)
}

func (t *GenericTree[T]) heightHelper(node *TreeNode[T]) int {
	if node == nil || len(node.Children) == 0 {
		return 0
	}
	
	maxHeight := 0
	for _, child := range node.Children {
		height := t.heightHelper(child)
		if height > maxHeight {
			maxHeight = height
		}
	}
	
	return maxHeight + 1
}

// Size returns the number of nodes
func (t *GenericTree[T]) Size() int {
	return t.size
}

// AddChild is a helper to add a child to a node
func (n *TreeNode[T]) AddChild(value T) *TreeNode[T] {
	child := &TreeNode[T]{
		Value:    value,
		Children: make([]*TreeNode[T], 0),
		Parent:   n,
	}
	n.Children = append(n.Children, child)
	return child
}

