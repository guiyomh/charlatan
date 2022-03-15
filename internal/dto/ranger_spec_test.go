//go:build spec || test

package dto_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/guiyomh/charlatan/internal/dto"
)

func TestSpecRanger(t *testing.T) {
	t.Run("Should parse a numeric range", func(t *testing.T) {
		id := dto.SetID("user_{1..5}")
		expectedList := []string{
			"user_1",
			"user_2",
			"user_3",
			"user_4",
			"user_5",
		}
		expectedIterators := []string{
			"1",
			"2",
			"3",
			"4",
			"5",
		}

		list, iterators, err := dto.MakeRange(id)
		if assert.NoError(t, err) {
			if assert.Len(t, list, 5) {
				assert.Equal(t, expectedList, list)
			}
			if assert.Len(t, iterators, 5) {
				assert.Equal(t, expectedIterators, iterators)
			}
		}
	})

	t.Run("Should parse a list range", func(t *testing.T) {
		id := dto.SetID("Admin_{bob,alice}")
		expectedList := []string{
			"Admin_bob",
			"Admin_alice",
		}
		expectedIterators := []string{
			"bob",
			"alice",
		}

		list, iterators, err := dto.MakeRange(id)
		if assert.NoError(t, err) {
			if assert.Len(t, list, 2) {
				assert.Equal(t, expectedList, list)
			}
			if assert.Len(t, iterators, 2) {
				assert.Equal(t, expectedIterators, iterators)
			}
		}
	})
}
