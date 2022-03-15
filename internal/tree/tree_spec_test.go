//go:build spec || test

package tree_test

import (
	"testing"

	"github.com/guiyomh/charlatan/internal/dto"
	"github.com/guiyomh/charlatan/internal/tree"
	"github.com/stretchr/testify/assert"
)

func TestSpecTree(t *testing.T) {

	var myTree tree.Tree

	root := dto.Row{
		Fields: dto.Fields{
			dto.Field("bob"): "@bob",
		},
		Meta: dto.Meta{
			RecordID: "root",
		},
	}

	bob := dto.Row{
		Fields: dto.Fields{
			dto.Field("foo"): "baz",
		},
		Meta: dto.Meta{
			RecordID: "bob",
		},
	}

	bobette := dto.Row{
		Fields: dto.Fields{
			dto.Field("bazou"): "bar",
		},
		Meta: dto.Meta{
			RecordID: "bobette",
		},
	}

	myTree.Add(tree.NewNode(root))
	myTree.Add(tree.NewNode(bob))
	myTree.Add(tree.NewNode(bobette))

	records := make([]string, 0)

	stringify := func(n *tree.Node, err error) error {
		if err != nil {
			return err
		}
		row := n.Value()
		records = append(records, string(row.Meta.RecordID))

		return nil
	}

	want := []string{"bob", "root", "bobette"}

	myTree.InOrderTraverse(stringify)

	assert.Equal(t, want, records)
}
