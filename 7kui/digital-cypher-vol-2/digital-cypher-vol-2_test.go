package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecode(t *testing.T) {
	{
		result := Decode([]int{20, 12, 18, 30, 21}, 1939)
		expect := "scout"
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := Decode([]int{14, 10, 22, 29, 6, 27, 19, 18, 6, 12, 8}, 1939)
		expect := "masterpiece"
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
}

/*
	Expect(Decode([]int{20, 12, 18, 30, 21}, 1939)).To(Equal("scout"))
	Expect(Decode([]int{14, 10, 22, 29, 6, 27, 19, 18, 6, 12, 8}, 1939)).To(Equal("masterpiece"))
*/
