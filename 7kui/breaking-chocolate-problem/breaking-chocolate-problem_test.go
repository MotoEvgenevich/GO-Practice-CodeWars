package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBreakChocolate(t *testing.T) {
	{
		result := BreakChocolate(5, 5)
		expect := 24
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := BreakChocolate(7, 4)
		expect := 27
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := BreakChocolate(1, 1)
		expect := 0
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := BreakChocolate(0, 0)
		expect := 0
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := BreakChocolate(6, 1)
		expect := 5
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
}

/*
   Expect(BreakChocolate(5,5)).To(Equal(24));
   Expect(BreakChocolate(7,4)).To(Equal(27));
   Expect(BreakChocolate(1,1)).To(Equal(0));
   Expect(BreakChocolate(0,0)).To(Equal(0), "What If I Told You There is No Chocolate?");
   Expect(BreakChocolate(6,1)).To(Equal(5));
*/
