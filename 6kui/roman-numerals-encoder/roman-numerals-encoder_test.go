package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	{
		result := Solution(1)
		expect := "I"
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := Solution(2)
		expect := "II"
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := Solution(3)
		expect := "III"
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := Solution(4)
		expect := "IV"
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := Solution(5)
		expect := "V"
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}

	{
		result := Solution(9)
		expect := "IX"
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := Solution(58)
		expect := "LVIII"
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := Solution(1994)
		expect := "MCMXCIV"
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}

	{
		result := Solution(0)
		expect := ""
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := Solution(182)
		expect := "CLXXXII"
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}

	{
		result := Solution(1990)
		expect := "MCMXC"
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := Solution(1875)
		expect := "MDCCCLXXV"
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
}

/*
   Expect(Solution(182)).To(Equal("CLXXXII"))
   Expect(Solution(1990)).To(Equal("MCMXC"))
   Expect(Solution(1875)).To(Equal("MDCCCLXXV"))
*/
