//go:build spec || test

package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/guiyomh/charlatan/internal/dto"
)

func TestSpecNodeHasDependency(t *testing.T) {
	n := Node{
		val: dto.Row{
			Fields: dto.Fields{
				dto.Field("bob"): "@bob",
			},
		},
		left: &Node{
			val: dto.Row{
				Fields: dto.Fields{
					dto.Field("foo"): "@foo",
					dto.Field("baz"): "@baz",
				},
			},
			left: &Node{
				val: dto.Row{
					Fields: dto.Fields{
						dto.Field("foo1"): "@foo1",
						dto.Field("baz1"): "@baz2",
					},
				},
			},
		},
		right: &Node{
			val: dto.Row{
				Fields: dto.Fields{
					dto.Field("bar"): "@bar",
				},
			},
		},
	}

	assert.True(t, n.HasDependencyOf("foo"))
	assert.True(t, n.HasDependencyOf("foo1"))
	assert.True(t, n.HasDependencyOf("baz2"))
	assert.True(t, n.HasDependencyOf("baz"))
	assert.True(t, n.HasDependencyOf("bar"))
	assert.True(t, n.HasDependencyOf("bob"))
	assert.False(t, n.HasDependencyOf("biloute"))
}

func TestSpecNodeInsert(t *testing.T) {
	node := &Node{
		val: dto.Row{
			Fields: dto.Fields{
				dto.Field("bob"): "@bob",
			},
		},
	}
	bob := &Node{
		val: dto.Row{
			Fields: dto.Fields{
				dto.Field("foo"): "baz",
			},
			Meta: dto.Meta{
				RecordID: "bob",
			},
		},
	}
	bobette := &Node{
		val: dto.Row{
			Fields: dto.Fields{
				dto.Field("bazou"): "bar",
			},
			Meta: dto.Meta{
				RecordID: "bobette",
			},
		},
	}
	if assert.Nil(t, node.left) {
		node.insert(bob)
		assert.Equal(t, bob, node.left)
	}
	if assert.Nil(t, node.right) {
		node.insert(bobette)
		assert.Equal(t, bobette, node.right)
	}

}
