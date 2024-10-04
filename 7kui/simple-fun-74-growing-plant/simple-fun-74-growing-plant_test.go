package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGrowingPlant(t *testing.T) {
	{
		result := GrowingPlant(100, 10, 910)
		var expect int = 10
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := GrowingPlant(10, 9, 4)
		var expect int = 1
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := GrowingPlant(5, 2, 5)
		var expect int = 1
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := GrowingPlant(5, 2, 6)
		var expect int = 2
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
}

/*
    It("GrowingPlant(100, 10, 910)", func() { Expect(GrowingPlant(100, 10, 910)).To(Equal(10)) })
    It("GrowingPlant(10, 9, 4)", func() { Expect(GrowingPlant(10, 9, 4)).To(Equal(1)) })
    It("GrowingPlant(5, 2, 5)", func() { Expect(GrowingPlant(5, 2, 5)).To(Equal(1)) })
    It("GrowingPlant(5, 2, 6)", func() { Expect(GrowingPlant(5, 2, 6)).To(Equal(2)) })
})

*/
