package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartsSums(t *testing.T) {
	{
		result := PartsSums([]uint64{})
		expect := []uint64{0}
		assert.Equal(t, expect, result, "result value doesn't equal expect value")
	}
	{
		result := PartsSums([]uint64{0, 1, 3, 6, 10})
		expect := []uint64{20, 20, 19, 16, 10, 0}
		assert.Equal(t, expect, result, "result value doesn't equal expect value")
	}
	{
		result := PartsSums([]uint64{1, 2, 3, 4, 5, 6})
		expect := []uint64{21, 20, 18, 15, 11, 6, 0}
		assert.Equal(t, expect, result, "result value doesn't equal expect value")
	}
	{
		result := PartsSums([]uint64{744125, 935, 407, 454, 430, 90, 144, 6710213, 889, 810, 2579358})
		expect := []uint64{10037855, 9293730, 9292795, 9292388, 9291934, 9291504, 9291414, 9291270, 2581057, 2580168, 2579358, 0}
		assert.Equal(t, expect, result, "result value doesn't equal expect value")
	}
}

/*
  Expect(PartsSums([]uint64{})).To(Equal([]uint64{0}))
       Expect(PartsSums([]uint64{0, 1, 3, 6, 10})).To(Equal([]uint64{20, 20, 19, 16, 10, 0}))
       Expect(PartsSums([]uint64{1, 2, 3, 4, 5, 6})).To(Equal([]uint64{21, 20, 18, 15, 11, 6, 0}))
       Expect(PartsSums([]uint64{744125, 935, 407, 454, 430, 90, 144, 6710213, 889, 810, 2579358})).To(Equal([]uint64{10037855, 9293730, 9292795, 9292388, 9291934, 9291504, 9291414, 9291270, 2581057, 2580168, 2579358, 00}))
     })
*/
