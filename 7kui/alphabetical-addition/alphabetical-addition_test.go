package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddLetters(t *testing.T) {
	{
		result := AddLetters([]rune{'a', 'b', 'c'})
		expect := 'f'
		assert.Equal(t, result, expect, "не соответвует result & expect")
	}
	{
		result := AddLetters([]rune{'a', 'b'})
		expect := 'c'
		assert.Equal(t, result, expect, "не соответвует result & expect")
	}
	{
		result := AddLetters([]rune{'z'})
		expect := 'z'
		assert.Equal(t, result, expect, "не соответвует result & expect")
	}
	{
		result := AddLetters([]rune{'z', 'a'})
		expect := 'a'
		assert.Equal(t, result, expect, "не соответвует result & expect")
	}
	{
		result := AddLetters([]rune{'y', 'c', 'b'})
		expect := 'd'
		assert.Equal(t, result, expect, "не соответвует result & expect")
	}
	{
		result := AddLetters([]rune{})
		expect := 'z'
		assert.Equal(t, result, expect, "не соответвует result & expect")
	}

}

/*
AddLetters([]rune{'a', 'b', 'c'}) = 'f'
AddLetters([]rune{'a', 'b'}) = 'c'
AddLetters([]rune{'z'}) = 'z'
AddLetters([]rune{'z', 'a'}) = 'a'
AddLetters([]rune{'y', 'c', 'b'}) = 'd' // notice the letters overflowing
AddLetters([]rune{}) = 'z'
*/
