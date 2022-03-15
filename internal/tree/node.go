package tree

import "github.com/guiyomh/charlatan/internal/dto"

// Node structure for Nodes
type Node struct {
	val   dto.Row
	left  *Node
	right *Node
}

// NewNode Returns a new pointer to an empty Node
func NewNode(val dto.Row) *Node {
	return &Node{val, nil, nil}
}

func (n *Node) Value() dto.Row {
	return n.val
}

func (n *Node) HasDependencyOf(other string) bool {

	if n.val.HasDependencyOf(other) {
		return true
	}

	if n.left != nil && n.left.HasDependencyOf(other) {
		return true
	}

	if n.right != nil && n.right.HasDependencyOf(other) {
		return true
	}

	return false
}

func (n *Node) insert(data *Node) {
	if n.HasDependencyOf(string(data.val.Meta.RecordID)) {
		if n.left == nil {
			n.left = data
		} else {
			n.left.insert(data)
		}

		return
	}
	if n.right == nil {
		n.right = data
	} else {
		n.right.insert(data)
	}
}

func inOrderTraverse(n *Node, f func(*Node, error) error, err error) error {
	if err != nil {
		return err
	}
	if n != nil {
		err = inOrderTraverse(n.left, f, err)
		err = f(n, err)
		err = inOrderTraverse(n.right, f, err)
		if err != nil {
			return err
		}
	}

	return nil
}
