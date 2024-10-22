package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFind(t *testing.T) {

	{
		result := FindUniq([]float32{1.0, 1.0, 1.0, 2.0, 1.0, 1.0})
		expect := float32(2)
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := FindUniq([]float32{0, 0, 0.55, 0, 0})
		expect := float32(0.55)
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
}

/*
   Expect(FindUniq([]float32{ 1.0, 1.0, 1.0, 2.0, 1.0, 1.0 })).To(Equal(float32(2)))
   Expect(FindUniq([]float32{ 0, 0, 0.55, 0, 0  })).To(Equal(float32(0.55)))
*/
