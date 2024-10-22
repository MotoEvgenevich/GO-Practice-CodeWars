package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	{
		result := Parse("iiisdoso")
		expect := []int{8, 64}
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := Parse("ooo")
		expect := []int{0, 0, 0}
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}

	{
		result := Parse("ioioio")
		expect := []int{1, 2, 3}
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := Parse("isoisoiso")
		expect := []int{1, 4, 25}
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}

	{
		result := Parse("codewars")
		expect := []int{0}
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := Parse("idoiido")
		expect := []int{0, 1}
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
}

/*
   It("just o", func() {
        Expect(Parse("ooo")).To(Equal([]int{0,0,0}))
    })
    It("o&i", func() {
        Expect(Parse("ioioio")).To(Equal([]int{1,2,3}))
    })
    It("o&i&d", func() {
        Expect(Parse("idoiido")).To(Equal([]int{0,1}))
    })
    It("o&i&d&s", func() {
        Expect(Parse("isoisoiso")).To(Equal([]int{1,4,25}))
    })
    It("codewars", func() {
        Expect(Parse("codewars")).To(Equal([]int{0}))
    })
*/
