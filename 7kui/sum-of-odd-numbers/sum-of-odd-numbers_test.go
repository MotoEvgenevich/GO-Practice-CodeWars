package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRowSumOddNumbers(t *testing.T) {
	{
		result := RowSumOddNumbers(1)
		expect := 1
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := RowSumOddNumbers(2)
		expect := 8
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := RowSumOddNumbers(13)
		expect := 2197
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := RowSumOddNumbers(19)
		expect := 6859
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := RowSumOddNumbers(41)
		expect := 68921
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := RowSumOddNumbers(42)
		expect := 74088
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := RowSumOddNumbers(74)
		expect := 405224
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := RowSumOddNumbers(86)
		expect := 636056
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := RowSumOddNumbers(93)
		expect := 804357
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := RowSumOddNumbers(101)
		expect := 1030301
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
}
func TestBestPractice(t *testing.T) {
	{
		result := bestPractice(1)
		expect := 1
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := bestPractice(2)
		expect := 8
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := bestPractice(13)
		expect := 2197
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := bestPractice(19)
		expect := 6859
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := bestPractice(41)
		expect := 68921
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := bestPractice(42)
		expect := 74088
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := bestPractice(74)
		expect := 405224
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := bestPractice(86)
		expect := 636056
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := bestPractice(93)
		expect := 804357
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := bestPractice(101)
		expect := 1030301
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
}

/*
    It("Testing for 1", func() { Expect(RowSumOddNumbers(1)).To(Equal(1)) })
    It("Testing for 2", func() { Expect(RowSumOddNumbers(2)).To(Equal(8)) })
    It("Testing for 13", func() { Expect(RowSumOddNumbers(13)).To(Equal(2197)) })
    It("Testing for 19", func() { Expect(RowSumOddNumbers(19)).To(Equal(6859)) })
    It("Testing for 41", func() { Expect(RowSumOddNumbers(41)).To(Equal(68921)) })
    It("Testing for 42", func() { Expect(RowSumOddNumbers(42)).To(Equal(74088)) })
    It("Testing for 74", func() { Expect(RowSumOddNumbers(74)).To(Equal(405224)) })
    It("Testing for 86", func() { Expect(RowSumOddNumbers(86)).To(Equal(636056)) })
    It("Testing for 93", func() { Expect(RowSumOddNumbers(93)).To(Equal(804357)) })
    It("Testing for 101", func() { Expect(RowSumOddNumbers(101)).To(Equal(1030301)) })
})
*/
