package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsTriangle(t *testing.T) {
	{
		result := IsTriangle(1, 2, 2)
		expect := true
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := IsTriangle(7, 2, 2)
		expect := false
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := IsTriangle(1, 2, 3)
		expect := false
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := IsTriangle(3, 1, 2)
		expect := false
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := IsTriangle(5, 1, 2)
		expect := false
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := IsTriangle(1, 2, 5)
		expect := false
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := IsTriangle(2, 5, 1)
		expect := false
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := IsTriangle(4, 2, 3)
		expect := true
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := IsTriangle(5, 1, 5)
		expect := true
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := IsTriangle(2, 2, 2)
		expect := true
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := IsTriangle(-1, 2, 3)
		expect := false
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := IsTriangle(1, -2, 3)
		expect := false
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := IsTriangle(1, 2, -3)
		expect := false
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := IsTriangle(0, 2, 3)
		expect := false
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}

}

/*
	It("works for some examples", func() {
		Expect(IsTriangle(1, 2, 2)).To(Equal(true))
		Expect(IsTriangle(7, 2, 2)).To(Equal(false))
		Expect(IsTriangle(1, 2, 3)).To(Equal(false))
		Expect(IsTriangle(3, 1, 2)).To(Equal(false))
		Expect(IsTriangle(5, 1, 2)).To(Equal(false))
		Expect(IsTriangle(1, 2, 5)).To(Equal(false))
		Expect(IsTriangle(2, 5, 1)).To(Equal(false))
		Expect(IsTriangle(4, 2, 3)).To(Equal(true))
		Expect(IsTriangle(5, 1, 5)).To(Equal(true))
		Expect(IsTriangle(2, 2, 2)).To(Equal(true))
		Expect(IsTriangle(-1, 2, 3)).To(Equal(false))
		Expect(IsTriangle(1, -2, 3)).To(Equal(false))
		Expect(IsTriangle(1, 2, -3)).To(Equal(false))
		Expect(IsTriangle(0, 2, 3.)).To(Equal(false))
*/
