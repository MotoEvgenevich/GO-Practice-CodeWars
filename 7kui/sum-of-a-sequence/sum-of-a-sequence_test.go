package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSequenceSum(t *testing.T) {
	{
		result := SequenceSum(2, 6, 2)
		expect := 12
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := SequenceSum(1, 5, 1)
		expect := 15
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := SequenceSum(1, 5, 3)
		expect := 5
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := SequenceSum(0, 15, 3)
		expect := 45
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := SequenceSum(16, 15, 3)
		expect := 0
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := SequenceSum(2, 24, 22)
		expect := 26
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := SequenceSum(2, 2, 2)
		expect := 2
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := SequenceSum(2, 2, 1)
		expect := 2
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := SequenceSum(1, 15, 3)
		expect := 35
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := SequenceSum(15, 1, 3)
		expect := 0
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
}

/*
   It("Sample tests", func() {
      dotest(2, 6, 2, 12)
      dotest(1, 5, 1, 15)
      dotest(1, 5, 3, 5)
      dotest(0, 15, 3, 45)
      dotest(16, 15, 3, 0)
      dotest(2, 24, 22, 26)
      dotest(2, 2, 2, 2)
      dotest(2, 2, 1, 2)
      dotest(1, 15, 3, 35)
      dotest(15, 1, 3, 0)
   })
*/
