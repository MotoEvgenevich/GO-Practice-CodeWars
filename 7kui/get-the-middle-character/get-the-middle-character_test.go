package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMiddle(t *testing.T) {
	{
		result := GetMiddle("test")
		expect := "es"
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := GetMiddle("testing")
		expect := "t"
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := GetMiddle("middle")
		expect := "dd"
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := GetMiddle("A")
		expect := "A"
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
}

/*
   It("Tests", func() {
     Expect(GetMiddle("test")).To(Equal("es"))
     Expect(GetMiddle("testing")).To(Equal("t"))
     Expect(GetMiddle("middle")).To(Equal("dd"))
     Expect(GetMiddle("A")).To(Equal("A"))
   })
})
*/
