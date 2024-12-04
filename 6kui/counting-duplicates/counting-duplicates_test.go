package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDuplicate_count(t *testing.T) {
	{
		result := Duplicate_count("abcde")
		expected := 0
		assert.Equal(t, expected, result, "не соотвествуют данным:  dotest(\"abcde\", 0)")
	}
	{
		result := Duplicate_count("abcdea")
		expected := 1
		assert.Equal(t, expected, result, "не соотвествуют данным:  dotest(\"abcdea\", 1)")
	}
	{
		result := Duplicate_count("abcdeaB11")
		expected := 3
		assert.Equal(t, expected, result, "не соотвествуют данным:  dotest(\"abcdeaB11\", 3)")
	}
	{
		result := Duplicate_count("indivisibility")
		expected := 1
		assert.Equal(t, expected, result, "не соотвествуют данным:  dotest(\"indivisibility\", 1)")
	}
}

/*
   dotest("abcde", 0)
   dotest("abcdea", 1)
   dotest("abcdeaB11", 3)
   dotest("indivisibility", 1)
*/
