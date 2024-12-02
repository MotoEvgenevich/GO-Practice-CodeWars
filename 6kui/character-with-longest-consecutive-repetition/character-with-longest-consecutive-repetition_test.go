package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestRepetition(t *testing.T) {
	{
		result := LongestRepetition("aaaabb")
		expect := Result{'a', 4}
		assert.Equal(t, expect, result, "result doesn't equal expect")
	}
	{
		result := LongestRepetition("bbbaaabaaaa")
		expect := Result{'a', 4}
		assert.Equal(t, expect, result, "result doesn't equal expect")
	}
	{
		result := LongestRepetition("cbdeuuu900")
		expect := Result{'u', 3}
		assert.Equal(t, expect, result, "result doesn't equal expect")
	}
	{
		result := LongestRepetition("abbbbb")
		expect := Result{'b', 5}
		assert.Equal(t, expect, result, "result doesn't equal expect")
	}
	{
		result := LongestRepetition("aabb")
		expect := Result{'a', 2}
		assert.Equal(t, expect, result, "result doesn't equal expect")
	}
	{
		result := LongestRepetition("ba")
		expect := Result{'b', 1}
		assert.Equal(t, expect, result, "result doesn't equal expect")
	}
	{
		result := LongestRepetition("")
		expect := Result{}
		assert.Equal(t, expect, result, "result doesn't equal expect")
	}
}

/*
	It("it should work with the sample tests", func() {
		Expect(LongestRepetition("aaaabb")).Should(Equal(Result{'a', 4}))
		Expect(LongestRepetition("bbbaaabaaaa")).Should(Equal(Result{'a', 4}))
		Expect(LongestRepetition("cbdeuuu900")).Should(Equal(Result{'u', 3}))
		Expect(LongestRepetition("abbbbb")).Should(Equal(Result{'b', 5}))
		Expect(LongestRepetition("aabb")).Should(Equal(Result{'a', 2}))
		Expect(LongestRepetition("ba")).Should(Equal(Result{'b', 1}))
		Expect(LongestRepetition("")).Should(Equal(Result{}))
*/
