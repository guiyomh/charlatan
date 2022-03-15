// Package tree provides tools to sort Row before persisting it
package tree

import (
	"sync"

	"github.com/guiyomh/charlatan/internal/dto"
)

// Tree Returns a binary search tree structure which contains only a root Node
type Tree struct {
	lock sync.RWMutex
	root *Node
}

// Add a value in the BSTree
func (t *Tree) Add(node *Node) {
	t.lock.Lock()
	defer t.lock.Unlock()
	if t.root == nil {
		t.root = node

		return
	}
	t.root.insert(node)
}

func (t *Tree) InOrderTraverse(f func(*Node, error) error) error {
	var err error
	t.lock.RLock()
	defer t.lock.RUnlock()

	return inOrderTraverse(t.root, f, err)
}

func ConvertTablesToTree(tables dto.Tables) *Tree {
	var t Tree

	for _, records := range tables {
		for _, row := range records {
			t.Add(NewNode(*row))
		}
	}

	return &t
}
