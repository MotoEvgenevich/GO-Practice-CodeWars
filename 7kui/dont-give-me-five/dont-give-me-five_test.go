package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDontGiveMeFive(t *testing.T) {
	{
		result := DontGiveMeFive(1, 9)
		expect := 8
		assert.Equal(t, expect, result, "result doesn't equal expect")
	}
	{
		result := DontGiveMeFive(4, 17)
		expect := 12
		assert.Equal(t, expect, result, "result doesn't equal expect")
	}
	{
		result := DontGiveMeFive(-9, -1)
		expect := 8
		assert.Equal(t, expect, result, "result doesn't equal expect")
	}
}

/*
 It("DontGiveMeFive(1, 9) should Equal 8", func() {
     Expect(DontGiveMeFive(1, 9)).To(Equal(8))
   })
   It("DontGiveMeFive(4, 17) should Equal 12", func() {
     Expect(DontGiveMeFive(4, 17)).To(Equal(12))
   })
	 DontGiveMeFive(-9, -1) should Equal 8
	 DontGiveMeFive(-55, 12) should Equal 56
	 (501, 599) should Equal 0
	 0, 60) should Equal 46
*/
