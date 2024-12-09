package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecode(t *testing.T) {
	{
		result := Decode("XXI")
		expect := 21
		assert.Equal(t, expect, result, "result doesn't equal expect")
	}
	{
		result := Decode("I")
		expect := 1
		assert.Equal(t, expect, result, "result doesn't equal expect")
	}
	{
		result := Decode("IV")
		expect := 4
		assert.Equal(t, expect, result, "result doesn't equal expect")
	}
	{
		result := Decode("MMVIII")
		expect := 2008
		assert.Equal(t, expect, result, "result doesn't equal expect")
	}
	{
		result := Decode("MDCLXVI")
		expect := 1666
		assert.Equal(t, expect, result, "result doesn't equal expect")
	}
}

/*
var _ = Describe("test roman to decimal converter", func() {
   It("should give decimal number from roman", func() {
     Expect(Decode("XXI")).To(Equal(21))

   })
   It("should give decimal number from roman", func() {
     Expect(Decode("I")).To(Equal(1))
   })
   It("should give decimal number from roman", func() {
     Expect(Decode("IV")).To(Equal(4))
   })
   It("should give decimal number from roman", func() {
     Expect(Decode("MMVIII")).To(Equal(2008))
   })
   It("should give decimal number from roman", func() {
     Expect(Decode("MDCLXVI")).To(Equal(1666))
*/
