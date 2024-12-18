package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLastDigit(t *testing.T) {

	{
		result := LastDigit([]int{0, 0})
		expect := 1
		assert.Equal(t, expect, result, "result value doesn't equal expect value")
	}

	{
		result := LastDigit([]int{1, 2})
		expect := 1
		assert.Equal(t, expect, result, "result value doesn't equal expect value")
	}
	{
		result := LastDigit([]int{3, 4, 5})
		expect := 1
		assert.Equal(t, expect, result, "result value doesn't equal expect value")
	}
	{
		result := LastDigit([]int{4, 3, 6})
		expect := 4
		assert.Equal(t, expect, result, "result value doesn't equal expect value")
	}
	{
		result := LastDigit([]int{7, 6, 21})
		expect := 1
		assert.Equal(t, expect, result, "result value doesn't equal expect value")
	}
	{
		result := LastDigit([]int{12, 30, 21})
		expect := 6
		assert.Equal(t, expect, result, "result value doesn't equal expect value")
	}
	{
		result := LastDigit([]int{2, 0, 1})
		expect := 1
		assert.Equal(t, expect, result, "result value doesn't equal expect value")
	}
	{
		result := LastDigit([]int{2, 2, 2, 0})
		expect := 4
		assert.Equal(t, expect, result, "result value doesn't equal expect value")
	}
	{
		result := LastDigit([]int{937640, 767456, 981242})
		expect := 0
		assert.Equal(t, expect, result, "result value doesn't equal expect value")
	}
	{
		result := LastDigit([]int{123232, 694022, 140249})
		expect := 6
		assert.Equal(t, expect, result, "result value doesn't equal expect value")
	}
	{
		result := LastDigit([]int{499942, 898102, 846073})
		expect := 6
		assert.Equal(t, expect, result, "result value doesn't equal expect value")
	}
}
func TestSingle(t *testing.T) {
	{
		result := LastDigit([]int{})
		expect := 1
		assert.Equal(t, expect, result, "result value doesn't equal expect value")
	}
	{
		result := LastDigit([]int{0, 0, 0})
		expect := 0
		assert.Equal(t, expect, result, "result value doesn't equal expect value")
	}
}

/*
	Expect(LastDigit([]int{})).To(Equal(1))
	Expect(LastDigit([]int{0, 0})).To(Equal(1))    // 0 ^ 0
	Expect(LastDigit([]int{0, 0, 0})).To(Equal(0)) // 0^(0 ^ 0) = 0^1 = 0
	Expect(LastDigit([]int{1, 2})).To(Equal(1))
	Expect(LastDigit([]int{3, 4, 5})).To(Equal(1))
	Expect(LastDigit([]int{4, 3, 6})).To(Equal(4))
	Expect(LastDigit([]int{7, 6, 21})).To(Equal(1))
	Expect(LastDigit([]int{12, 30, 21})).To(Equal(6))
	Expect(LastDigit([]int{2, 0, 1})).To(Equal(1))
	Expect(LastDigit([]int{2, 2, 2, 0})).To(Equal(4))
	Expect(LastDigit([]int{937640, 767456, 981242})).To(Equal(0))
	Expect(LastDigit([]int{123232, 694022, 140249})).To(Equal(6))
	Expect(LastDigit([]int{499942, 898102, 846073})).To(Equal(6))
*/
