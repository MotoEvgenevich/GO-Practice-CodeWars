package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEquableTriangle(t *testing.T) {
	{
		result := EquableTriangle(5, 12, 13)
		expect := true
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := EquableTriangle(2, 3, 4)
		expect := false
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
}

/*
  It("should return true", func() {
    Expect(EquableTriangle(5, 12, 13)).To(Equal(true))
  })
  It("should return false", func() {
    Expect(EquableTriangle(2, 3, 4)).To(Equal(false))
  })
*/
