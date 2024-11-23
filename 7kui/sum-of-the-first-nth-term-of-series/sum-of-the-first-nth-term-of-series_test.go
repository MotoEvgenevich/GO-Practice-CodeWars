package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXxx(t *testing.T) {
	{
		result := SeriesSum(1)
		expect := "1.00"
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := SeriesSum(2)
		expect := "1.25"
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := SeriesSum(3)
		expect := "1.39"
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := SeriesSum(4)
		expect := "1.49"
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
}

/*
   Expect(SeriesSum(1)).To(Equal("1.00"), "n = 1")
   Expect(SeriesSum(2)).To(Equal("1.25"), "n = 2")
   Expect(SeriesSum(3)).To(Equal("1.39"), "n = 3")
   Expect(SeriesSum(4)).To(Equal("1.49"), "n = 4")
*/
