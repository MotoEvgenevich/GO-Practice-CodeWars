package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseWords(t *testing.T) {
	{
		result := ReverseWords("The quick brown fox jumps over the lazy dog.")
		expect := "ehT kciuq nworb xof spmuj revo eht yzal .god"
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := ReverseWords("apple")
		expect := "elppa"
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := ReverseWords("a b c d")
		expect := "a b c d"
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := ReverseWords("double  spaced  words")
		expect := "elbuod  decaps  sdrow"
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := ReverseWords("stressed desserts")
		expect := "desserts stressed"
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := ReverseWords("hello hello")
		expect := "olleh olleh"
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
}

/*
   Expect(ReverseWords("The quick brown fox jumps over the lazy dog.")).To(Equal("ehT kciuq nworb xof spmuj revo eht yzal .god"))
   Expect(ReverseWords("apple")).To(Equal("elppa"))
   Expect(ReverseWords("a b c d")).To(Equal("a b c d"))
   Expect(ReverseWords("double  spaced  words")).To(Equal("elbuod  decaps  sdrow"))
   Expect(ReverseWords("stressed desserts")).To(Equal("desserts stressed"))
   Expect(ReverseWords("hello hello")).To(Equal("olleh olleh"))
*/
