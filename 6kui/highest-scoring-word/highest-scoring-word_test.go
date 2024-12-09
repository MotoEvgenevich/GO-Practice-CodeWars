package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHigh(t *testing.T) {
	{
		actual := High("man i need a taxi up to ubud")
		expect := "taxi"
		assert.Equal(t, expect, actual, "result doesn't equal expect")
	}
	{
		actual := High("what time are we climbing up the volcano")
		expect := "volcano"
		assert.Equal(t, expect, actual, "result doesn't equal expect")
	}
	{
		actual := High("take me to semynak")
		expect := "semynak"
		assert.Equal(t, expect, actual, "result doesn't equal expect")
	}
	{
		actual := High("aa b")
		expect := "aa"
		assert.Equal(t, expect, actual, "result doesn't equal expect")
	}
	{
		actual := High("b aa")
		expect := "b"
		assert.Equal(t, expect, actual, "result doesn't equal expect")
	}
	{
		actual := High("bb d")
		expect := "bb"
		assert.Equal(t, expect, actual, "result doesn't equal expect")
	}
	{
		actual := High("d bb")
		expect := "d"
		assert.Equal(t, expect, actual, "result doesn't equal expect")
	}
	{
		actual := High("aaa b")
		expect := "aaa"
		assert.Equal(t, expect, actual, "result doesn't equal expect")
	}

}

/*
   {"man i need a taxi up to ubud","taxi"},
   {"what time are we climbing up the volcano","volcano"},
   {"take me to semynak","semynak"},
   {"aa b", "aa"},
   {"b aa", "b"},
   {"bb d", "bb"},
   {"d bb", "d"},
   {"aaa b", "aaa"},
*/
