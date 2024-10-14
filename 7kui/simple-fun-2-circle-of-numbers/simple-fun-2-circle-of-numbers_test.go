package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCircleOfNumbers(t *testing.T) {
	{
		result := CircleOfNumbers(10, 2)
		expect := 7
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := CircleOfNumbers(10, 7)
		expect := 2
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := CircleOfNumbers(4, 1)
		expect := 3
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := CircleOfNumbers(6, 3)
		expect := 0
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := CircleOfNumbers(20, 0)
		expect := 10
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
}

/*
var _ = Describe("Basic Tests", func() {
    It("CircleOfNumbers(10, 2)", func() { Expect(CircleOfNumbers(10, 2)).To(Equal(7)) })
    It("CircleOfNumbers(10, 7)", func() { Expect(CircleOfNumbers(10, 7)).To(Equal(2)) })
    It("CircleOfNumbers(4, 1)", func() { Expect(CircleOfNumbers(4, 1)).To(Equal(3)) })
    It("CircleOfNumbers(6, 3)", func() { Expect(CircleOfNumbers(6, 3)).To(Equal(0)) })
    It("CircleOfNumbers(20, 0)", func() { Expect(CircleOfNumbers(20, 0)).To(Equal(10)) })
*/
