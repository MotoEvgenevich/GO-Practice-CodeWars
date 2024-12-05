package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPrime(t *testing.T) {
	{
		result := IsPrime(0)
		expect := false
		assert.Equal(t, expect, result, "result doesn't equal expect")
	}
	{
		result := IsPrime(1)
		expect := false
		assert.Equal(t, expect, result, "result doesn't equal expect")
	}
	{
		result := IsPrime(2)
		expect := true
		assert.Equal(t, expect, result, "result doesn't equal expect")
	}
	{
		result := IsPrime(73)
		expect := true
		assert.Equal(t, expect, result, "result doesn't equal expect")
	}
	{
		result := IsPrime(75)
		expect := false
		assert.Equal(t, expect, result, "result doesn't equal expect")
	}
	{
		result := IsPrime(-1)
		expect := false
		assert.Equal(t, expect, result, "result doesn't equal expect")
	}
	{
		result := IsPrime(3)
		expect := true
		assert.Equal(t, expect, result, "result doesn't equal expect")
	}
	{
		result := IsPrime(5)
		expect := true
		assert.Equal(t, expect, result, "result doesn't equal expect")
	}
}

/*
   It("Basic tests", func() {
     doTest(0, false)
     doTest(1, false)
     doTest(2, true)
     doTest(73, true)
     doTest(75, false)
     doTest(-1, false)
   })

   It("Test prime", func() {
     doTest(3, true)
     doTest(5, true)
     doTest(7, true)
     doTest(41, true)
     doTest(5099, true)
   })

   It("Test not prime", func() {
     doTest(4, false)
     doTest(6, false)
     doTest(8, false)
     doTest(9, false)
     doTest(45, false)
     doTest(-5, false)
     doTest(-8, false)
     doTest(-41, false)
   })
})
*/
