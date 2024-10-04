package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	{
		result := Add(1)(3)
		var expect int = 4
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := Add(1)(-5)
		var expect int = -4
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := Add(3)(20)
		var expect int = 23
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
}

/*
  It("basic tests", func() {
    Expect(Add(1)(3)).To(Equal(4))
    Expect(Add(1)(-5)).To(Equal(-4))
    Expect(Add(3)(20)).To(Equal(23))
  })
  It("make sure Add() is pure", func() {
    addThree := Add(3)
    Expect(addThree(5)).To(Equal(8))
    _ = Add(4)
    Expect(addThree(5)).To(Equal(8))
  })
  It("random calculations", func() {
    for i := 0; i < 100; i++ {
      num1 := rand.Intn(1001) - 500
      num2 := rand.Intn(1001) - 500
      Expect(Add(num1)(num2)).To(Equal(num1 + num2))
    }
*/
