package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleMultiplication(t *testing.T) {
	{
		result := SimpleMultiplication(1)
		expect := 9
		assert.Equal(t, expect, result, "result value doesn't equal expect value")
	}
	{
		result := SimpleMultiplication(3)
		expect := 27
		assert.Equal(t, expect, result, "result value doesn't equal expect value")
	}
	{
		result := SimpleMultiplication(21)
		expect := 189
		assert.Equal(t, expect, result, "result value doesn't equal expect value")
	}
	{
		result := SimpleMultiplication(23)
		expect := 207
		assert.Equal(t, expect, result, "result value doesn't equal expect value")
	}
	{
		result := SimpleMultiplication(2)
		expect := 16
		assert.Equal(t, expect, result, "result value doesn't equal expect value")
	}
	{
		result := SimpleMultiplication(4)
		expect := 32
		assert.Equal(t, expect, result, "result value doesn't equal expect value")
	}
	{
		result := SimpleMultiplication(22)
		expect := 176
		assert.Equal(t, expect, result, "result value doesn't equal expect value")
	}
	{
		result := SimpleMultiplication(26)
		expect := 208
		assert.Equal(t, expect, result, "result value doesn't equal expect value")
	}
}

/*
  It("Odd inputs", func() {
      DoTest(1, 9)
      DoTest(3, 27)
      DoTest(21, 189)
      DoTest(23, 207)
  })

  It("Even inputs", func() {
      DoTest(2, 16)
      DoTest(4, 32)
      DoTest(22, 176)
      DoTest(26, 208)
  })
*/
