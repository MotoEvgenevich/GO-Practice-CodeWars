package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDigits(t *testing.T) {
	{
		result := Digits(12345)
		expect := 5
		assert.Equal(t, result, expect, "result doesn't equal expect")

	}
	{
		result := Digits(9876543210)
		expect := 10
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}

	{
		result := Digits(0)
		expect := 1
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := Digits(5)
		expect := 1
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
}

/*
   It("Sample tests", func() {
      dotest(5, 1)
      dotest(12345, 5)
      dotest(9876543210, 10)
   })
*/
