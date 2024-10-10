package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRangeBitCount(t *testing.T) {

	{
		result := RangeBitCount(2, 7)
		expect := 11
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}

	{
		result := RangeBitCount(0, 1)
		expect := 1
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}

	{
		result := RangeBitCount(4, 4)
		expect := 1
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
}

/*
   It("Testing 2 and 7",func() {Expect(RangeBitCount(2,7)).To(Equal(11))})
   It("Testing 0 and 1",func() {Expect(RangeBitCount(0,1)).To(Equal(1))})
   It("Testing 4 and 4",func() {Expect(RangeBitCount(4,4)).To(Equal(1))})
*/
