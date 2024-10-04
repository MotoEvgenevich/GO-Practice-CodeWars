package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtraPerfect(t *testing.T) {

	{
		result := ExtraPerfect(3)
		expect := []int{1, 3}
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := ExtraPerfect(5)
		expect := []int{1, 3, 5}
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := ExtraPerfect(7)
		expect := []int{1, 3, 5, 7}
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := ExtraPerfect(28)
		expect := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23, 25, 27}
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := ExtraPerfect(39)
		expect := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23, 25, 27, 29, 31, 33, 35, 37, 39}
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
}
func TestBestPractice(t *testing.T) {
	{
		result := BestPractice(3)
		expect := []int{1, 3}
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := BestPractice(5)
		expect := []int{1, 3, 5}
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := BestPractice(7)
		expect := []int{1, 3, 5, 7}
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := BestPractice(28)
		expect := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23, 25, 27}
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := BestPractice(39)
		expect := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23, 25, 27, 29, 31, 33, 35, 37, 39}
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
}

/*
  It("ExtraPerfect(3)", func() { Expect(ExtraPerfect(3)).To(Equal([]int{1, 3})) })
    It("ExtraPerfect(5)", func() { Expect(ExtraPerfect(5)).To(Equal([]int{1, 3, 5})) })
    It("ExtraPerfect(7)", func() { Expect(ExtraPerfect(7)).To(Equal([]int{1, 3, 5, 7})) })
    It("ExtraPerfect(28)", func() { Expect(ExtraPerfect(28)).To(Equal([]int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23, 25, 27})) })
    It("ExtraPerfect(39)", func() { Expect(ExtraPerfect(39)).To(Equal([]int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23, 25, 27, 29, 31, 33, 35, 37, 39})) })
})
*/
