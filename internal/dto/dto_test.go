//go:build !spec || test

package dto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApplyTemplate(t *testing.T) {

	t.Run("Should apply one template", func(t *testing.T) {
		f := Fixture{
			SetID("user_{1..3} (extends myTpl)"): Set{
				Field("email"): EntryValue("foo@bar.com"),
			},
			SetID("myTpl (template)"): Set{
				Field("first_name"): EntryValue("john"),
				Field("last_name"):  EntryValue("doe"),
			},
		}
		expectedFixture := Fixture{
			SetID("user_{1..3}"): Set{
				Field("email"):      EntryValue("foo@bar.com"),
				Field("first_name"): EntryValue("john"),
				Field("last_name"):  EntryValue("doe"),
			},
		}

		assert.Equal(t, expectedFixture, applyTemplate(f))
	})

	t.Run("Should erase template field by Set field", func(t *testing.T) {
		f := Fixture{
			SetID("user_{1..3} (extends myTpl)"): Set{
				Field("first_name"): EntryValue("barry"),
				Field("email"):      EntryValue("foo@bar.com"),
			},
			SetID("myTpl (template)"): Set{
				Field("first_name"): EntryValue("john"),
				Field("last_name"):  EntryValue("doe"),
			},
		}
		expectedFixture := Fixture{
			SetID("user_{1..3}"): Set{
				Field("email"):      EntryValue("foo@bar.com"),
				Field("first_name"): EntryValue("barry"),
				Field("last_name"):  EntryValue("doe"),
			},
		}

		assert.Equal(t, expectedFixture, applyTemplate(f))
	})

	t.Run("Should apply many templates", func(t *testing.T) {
		f := Fixture{
			SetID("user_{1..3} (extends myTpl)"): Set{
				Field("email"): EntryValue("foo@bar.com"),
			},
			SetID("admin_{4..6} (extends otherTpl)"): Set{
				Field("email"): EntryValue("admin@bar.com"),
			},
			SetID("myTpl (template)"): Set{
				Field("first_name"): EntryValue("john"),
				Field("last_name"):  EntryValue("doe"),
			},
			SetID("otherTpl (template)"): Set{
				Field("is_admin"): EntryValue(true),
			},
		}
		expectedFixture := Fixture{
			SetID("user_{1..3}"): Set{
				Field("email"):      EntryValue("foo@bar.com"),
				Field("first_name"): EntryValue("john"),
				Field("last_name"):  EntryValue("doe"),
			},
			SetID("admin_{4..6}"): Set{
				Field("email"):    EntryValue("admin@bar.com"),
				Field("is_admin"): EntryValue(true),
			},
		}

		assert.Equal(t, expectedFixture, applyTemplate(f))
	})

}

func TestMergeSet(t *testing.T) {

	t.Run("Should merge two Set into one", func(t *testing.T) {
		set1 := Set{
			Field("first_name"): EntryValue("john"),
		}
		set2 := Set{
			Field("last_name"): EntryValue("doe"),
		}
		expectSet := Set{
			Field("first_name"): EntryValue("john"),
			Field("last_name"):  EntryValue("doe"),
		}

		assert.Equal(t, expectSet, mergeSet(set1, set2))
	})
}

func TestReplaceCurrent(t *testing.T) {
	t.Run("Should replace {current}", func(t *testing.T) {
		iteration := "1"
		row := Row{
			Fields: Fields{
				Field("foo"): "@bar_{current}",
				Field("baz"): "@biloute_{current}",
			},
		}
		expectedRow := Row{
			Fields: Fields{
				Field("foo"): "@bar_1",
				Field("baz"): "@biloute_1",
			},
		}

		result := replaceCurrent(iteration, &row)

		assert.Equal(t, expectedRow, *result)
	})
}

func TestRowWHasDependancies(t *testing.T) {
	t.Run("Should has dependencies", func(t *testing.T) {
		row := Row{
			Fields: Fields{
				Field("foo"): "@bar_1",
				Field("baz"): "bar",
			},
		}
		assert.True(t, row.HasDependencies())
	})
	t.Run("Should not have any dependencies", func(t *testing.T) {
		row := Row{
			Fields: Fields{
				Field("foo"): "foofoo",
				Field("baz"): "bar",
			},
		}
		assert.False(t, row.HasDependencies())
	})
}

func TestRowWHasDependancyOf(t *testing.T) {
	t.Run("Should has dependency with a specific row", func(t *testing.T) {
		row := Row{
			Fields: Fields{
				Field("foo"): "@bar_1",
				Field("baz"): "bar",
			},
		}
		assert.True(t, row.HasDependencyOf("bar_1"))
	})
	t.Run("Should not have any dependency with a specific row", func(t *testing.T) {
		row := Row{
			Fields: Fields{
				Field("foo"): "@bar_1",
				Field("baz"): "bar",
			},
		}
		assert.False(t, row.HasDependencyOf("@foo_1"))
	})
}
