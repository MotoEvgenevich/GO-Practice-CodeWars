package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountDivisors(t *testing.T) {
	{
		result := Divisors(1)
		expect := 1
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := Divisors(10)
		expect := 4
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := Divisors(11)
		expect := 2
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := Divisors(54)
		expect := 8
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := Divisors(64)
		expect := 7
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
}

/*
)
var _ = Describe("Should pass some basic tests", func() {
	It("Divisors(1)", func() {Expect(Divisors(1)).To(Equal(1))})
	It("Divisors(10)", func() {Expect(Divisors(10)).To(Equal(4))})
	It("Divisors(11)", func() {Expect(Divisors(11)).To(Equal(2))})
	It("Divisors(54)", func() {Expect(Divisors(54)).To(Equal(8))})
	It("Divisors(64)", func() {Expect(Divisors(64)).To(Equal(7))})
})
*/
