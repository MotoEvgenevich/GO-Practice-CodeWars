package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumberOfPairs(t *testing.T) {
	{
		result := NumberOfPairs([]string{"red", "red"})
		expect := 1
		assert.Equal(t, expect, result, "result value doesn't equal expect value")
	}
	{
		result := NumberOfPairs([]string{"red", "green", "blue"})
		expect := 0
		assert.Equal(t, expect, result, "result value doesn't equal expect value")
	}
	{
		result := NumberOfPairs([]string{"gray", "black", "purple", "purple", "gray", "black"})
		expect := 3
		assert.Equal(t, expect, result, "result value doesn't equal expect value")
	}
	{
		result := NumberOfPairs([]string{})
		expect := 0
		assert.Equal(t, expect, result, "result value doesn't equal expect value")
	}
	{
		result := NumberOfPairs([]string{"red", "green", "blue", "blue", "red", "green", "red", "red", "red"})
		expect := 4
		assert.Equal(t, expect, result, "result value doesn't equal expect value")
	}
}

/*
Expect(NumberOfPairs([]string{"red", "red"})).To(Equal(1))
		Expect(NumberOfPairs([]string{"red", "green", "blue"})).To(Equal(0))
		Expect(NumberOfPairs([]string{"gray", "black", "purple", "purple", "gray", "black"})).To(Equal(3))
		Expect(NumberOfPairs([]string{})).To(Equal(0))
		Expect(NumberOfPairs([]string{"red", "green", "blue", "blue", "red", "green", "red", "red", "red"})).To(Equal(4))
*/
