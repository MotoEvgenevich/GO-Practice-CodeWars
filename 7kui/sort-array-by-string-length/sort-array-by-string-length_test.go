package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortByLength(t *testing.T) {
	{
		expect := []string{"i", "to", "beg", "life"}
		result := SortByLength([]string{"beg", "life", "i", "to"})
		assert.Equal(t, expect, result, "expect doesn't equal result")
	}
	{
		expect := []string{"", "pizza", "brains", "moderately"}
		result := SortByLength([]string{"", "moderately", "brains", "pizza"})
		assert.Equal(t, expect, result, "expect doesn't equal result")
	}
	{
		expect := []string{"short", "longer", "longest"}
		result := SortByLength([]string{"longer", "longest", "short"})
		assert.Equal(t, expect, result, "expect doesn't equal result")
	}
	{
		expect := []string{"a", "of", "dog", "food"}
		result := SortByLength([]string{"dog", "food", "a", "of"})
		assert.Equal(t, expect, result, "expect doesn't equal result")
	}
	{
		expect := []string{"", "bees", "eloquent", "dictionary"}
		result := SortByLength([]string{"", "dictionary", "eloquent", "bees"})
		assert.Equal(t, expect, result, "expect doesn't equal result")
	}
	{
		expect := []string{"a short sentence", "a longer sentence", "the longest sentence"}
		result := SortByLength([]string{"a longer sentence", "the longest sentence", "a short sentence"})
		assert.Equal(t, expect, result, "expect doesn't equal result")
	}
}
