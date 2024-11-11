package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncode(t *testing.T) {
	{
		result := Encode("scout", 1939)
		expect := []int{20, 12, 18, 30, 21}
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := Encode("masterpiece", 1939)
		expect := []int{14, 10, 22, 29, 6, 27, 19, 18, 6, 12, 8}
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
}

/*
	Expect(Encode("scout", 1939)).To(Equal([]int{20, 12, 18, 30, 21}))
	Expect(Encode("masterpiece", 1939)).To(Equal([]int{14, 10, 22, 29, 6, 27, 19, 18, 6, 12, 8}))
*/
