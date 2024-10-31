package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasUniqueChar(t *testing.T) {
	{
		result := HasUniqueChar("  nAa")
		expect := false
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := HasUniqueChar("abcde")
		expect := true
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := HasUniqueChar("++-")
		expect := false
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := HasUniqueChar("AaBbC")
		expect := true
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
}

/*
   Expect(HasUniqueChar("  nAa")).To(Equal(false))
   Expect(HasUniqueChar("abcde")).To(Equal(true))
   Expect(HasUniqueChar("++-")).To(Equal(false))
   Expect(HasUniqueChar("AaBbC")).To(Equal(true))
*/
