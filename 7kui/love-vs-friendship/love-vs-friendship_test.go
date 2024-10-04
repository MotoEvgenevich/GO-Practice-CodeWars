package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWordsToMarks(t *testing.T) {
	{
		result := WordsToMarks("attitude")
		expect := 100
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := WordsToMarks("friends")
		expect := 75
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := WordsToMarks("family")
		expect := 66
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := WordsToMarks("selfness")
		expect := 99
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := WordsToMarks("knowledge")
		expect := 96
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
}

/*
    It("Testing for attitude", func() { Expect(WordsToMarks("attitude")).To(Equal(100)) })
    It("Testing for friends", func() { Expect(WordsToMarks("friends")).To(Equal(75)) })
    It("Testing for family", func() { Expect(WordsToMarks("family")).To(Equal(66)) })
    It("Testing for selfness", func() { Expect(WordsToMarks("selfness")).To(Equal(99)) })
    It("Testing for knowledge", func() { Expect(WordsToMarks("knowledge")).To(Equal(96)) })
})

*/
