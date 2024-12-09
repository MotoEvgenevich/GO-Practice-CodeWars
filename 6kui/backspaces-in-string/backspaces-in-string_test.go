package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCleanString(t *testing.T) {
	{
		result := CleanString("abc#d##c")
		expect := "ac"
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := CleanString("abc####d##c#")
		expect := ""
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := CleanString("")
		expect := ""
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := CleanString("#######")
		expect := ""
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
}

/*
   Expect(CleanString("abc#d##c")).To(Equal("ac"))
   Expect(CleanString("abc####d##c#")).To(Equal(""))
   Expect(CleanString("")).To(Equal(""))
   Expect(CleanString("#######")).To(Equal(""))
*/
