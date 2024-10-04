package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScrabbleScore(t *testing.T) {

	{
		result := ScrabbleScore("")
		expect := 0
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := ScrabbleScore("a")
		expect := 1
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := ScrabbleScore("street")
		expect := 6
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := ScrabbleScore("STREET")
		expect := 6
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := ScrabbleScore(" a")
		expect := 1
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := ScrabbleScore("st re et")
		expect := 6
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := ScrabbleScore("f")
		expect := 4
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := ScrabbleScore("quirky")
		expect := 22
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := ScrabbleScore("MULTIBILLIONAIRE")
		expect := 20
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := ScrabbleScore("alacrity")
		expect := 13
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
}

/*
   Expect(ScrabbleScore("")).To(Equal(0))
   Expect(ScrabbleScore("a")).To(Equal(1))
   Expect(ScrabbleScore("street")).To(Equal(6))
   Expect(ScrabbleScore("STREET")).To(Equal(6))
   Expect(ScrabbleScore(" a")).To(Equal(1))
   Expect(ScrabbleScore("st re et")).To(Equal(6))
   Expect(ScrabbleScore("f")).To(Equal(4))
   Expect(ScrabbleScore("quirky")).To(Equal(22))
   Expect(ScrabbleScore("MULTIBILLIONAIRE")).To(Equal(20))
   Expect(ScrabbleScore("alacrity")).To(Equal(13))
*/
