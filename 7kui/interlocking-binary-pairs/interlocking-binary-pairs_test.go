package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInterlockable(t *testing.T) {
	{
		result := Interlockable(9, 4)
		expect := true
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := Interlockable(3, 6)
		expect := false
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := Interlockable(2, 5)
		expect := true
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := Interlockable(7, 1)
		expect := false
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := Interlockable(0, 8)
		expect := true
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}

}

func TestInterlockable2(t *testing.T) {
	{
		result := InterlockableAsString(9, 4)
		expect := true
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := InterlockableAsString(3, 6)
		expect := false
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := InterlockableAsString(2, 5)
		expect := true
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := InterlockableAsString(7, 1)
		expect := false
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := InterlockableAsString(0, 8)
		expect := true
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
}

/*
   dotest(9, 4, true)
   dotest(3, 6, false)
   dotest(2, 5, true)
   dotest(7, 1, false)
   dotest(0, 8, true)
*/
