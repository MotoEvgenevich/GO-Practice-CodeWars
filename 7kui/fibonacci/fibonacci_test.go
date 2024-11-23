package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFib(t *testing.T) {
	{
		result := Fib(1)
		expect := 1
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := Fib(2)
		expect := 1
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := Fib(3)
		expect := 2
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := Fib(4)
		expect := 3
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
}

/*
   Expect(Fib(1)).To(Equal(1))
   Expect(Fib(2)).To(Equal(1))
   Expect(Fib(3)).To(Equal(2))
   Expect(Fib(4)).To(Equal(3))
   Expect(Fib(5)).To(Equal(5))
*/
