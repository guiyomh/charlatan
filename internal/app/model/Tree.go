package model

// Node is a Node in a Binary Tree
type Node struct {
	Left  *Node
	Right *Node
	Key   string
	Value *Row
}

// NewNode Return a new empty node
func NewNode(row *Row) *Node {
	return &Node{Left: nil, Right: nil, Value: row, Key: row.Name}
}

// HasDependencyOf is True if the Row have a dependancy with the other row pass
func (me *Node) HasDependencyOf(name string) bool {
	if !me.Value.HasDependencies() {
		return false
	}
	for _, v := range me.Value.DependencyReference {
		if v.RecordName == name {
			return true
		}
	}
	return false
}

// LessThan return true if this row hasn't dependancy with other.
func (me *Node) LessThan(other *Node) bool {
	return !me.Value.HasDependencyOf(other.Value.Name)
}

// EqualTo return true if this row is equal to the other.
func (me *Node) EqualTo(other *Node) bool {
	return me.Value.Name == other.Value.Name
}

// GreaterThan return true if this row has dependancy of other.
func (me *Node) GreaterThan(other *Node) bool {
	return me.Value.HasDependencyOf(other.Value.Name)
}

// Add an existing node to this node's subtree
func (me *Node) Add(node *Node) *Node {
	if node.LessThan(me) {
		if me.Left == nil {
			me.Left = node
			return node
		}
		return me.Left.Add(node)
	} else {
		if me.Right == nil {
			me.Right = node
			return node
		}
		return me.Right.Add(node)
	}
}

// Minimum Return the left-most (smallest key) node in this node's subtree
func (me *Node) Minimum() *Node {
	for {
		if me.Left == nil {
			return me
		}
		me = me.Left
	}
}

// Maximum Return the right-most (largest key) node in this node's subtree
func (me *Node) Maximum() *Node {
	for {
		if me.Right == nil {
			return me
		}
		me = me.Right
	}
}

// WalkForward Call iterator for each node in this node's subtree in order, low to high
func (me *Node) WalkForward(iterator func(me *Node)) {
	if me.Left != nil {
		me.Left.WalkForward(iterator)
	}
	iterator(me)
	if me.Right != nil {
		me.Right.WalkForward(iterator)
	}
}

// WalkBackward Call iterator for each node in this node's subtree in reverse order, high to low
func (me *Node) WalkBackward(iterator func(me *Node)) {
	if me.Right != nil {
		me.Right.WalkBackward(iterator)
	}
	iterator(me)
	if me.Left != nil {
		me.Left.WalkBackward(iterator)
	}
}

// Tree is a binary tree of record
// It sort record with relation
type Tree struct {
	Records map[string]*Node
	Root    *Node
}

// Iterator is a func that can iterate a tree
type Iterator func(key string, value *Row)

// NewTree create an instance of tre
func NewTree() *Tree {
	return &Tree{Root: nil, Records: make(map[string]*Node, 0)}
}

// Add a record in the tree
func (me *Tree) Add(record *Row) {
	node := NewNode(record)
	me.Records[node.Value.Name] = node
	if me.Root != nil {
		me.Root.Add(node)
		return
	}
	me.Root = node

}

// Find and return the node with the supplied key in this subtree. Return nil if not found.
func (me *Tree) Find(key string) *Node {
	for k, node := range me.Records {
		if k == key {
			return node
		}
	}
	return nil
}

// First Return the first (lowest) key and value in the tree, or nil, nil if the tree is empty.
func (me *Tree) First() (string, *Row) {
	if me.Root == nil {
		return "", nil
	}
	node := me.Root.Minimum()
	return node.Key, node.Value
}

// Last Return the last (highest) key and value in the tree, or nil, nil if the tree is empty.
func (me *Tree) Last() (string, *Row) {
	if me.Root == nil {
		return "", nil
	}
	node := me.Root.Maximum()
	return node.Key, node.Value
}

// Walk Iterate the tree with the function in the supplied direction
func (me *Tree) Walk(iterator Iterator, forward bool) {
	if me.Root == nil {
		return
	}
	if forward {
		me.Root.WalkForward(func(node *Node) { iterator(node.Key, node.Value) })
	} else {
		me.Root.WalkBackward(func(node *Node) { iterator(node.Key, node.Value) })
	}
}
