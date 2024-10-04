package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumOfCubes(t *testing.T) {
	{
		result := SumCubes(10)
		expect := 3025
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := SumCubes(15)
		expect := 14400
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := SumCubes(25)
		expect := 105625
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := SumCubes(30)
		expect := 216225
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := SumCubes(0)
		expect := 0
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
}

/*
It("Here we go...",func() {
		fixed := [...]int{1,2,3,4,10,123}
		xpect := [...]int{1,9,36,100,3025,58155876}
*/
