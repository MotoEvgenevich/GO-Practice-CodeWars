package kata

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsNegativeZero(t *testing.T) {
	{
		result := IsNegativeZero(5.0)
		expect := false
		assert.Equal(t, result, expect, "doesn't match result & expect")
	}
	{
		result := IsNegativeZero(4.0)
		expect := false
		assert.Equal(t, result, expect, "doesn't match result & expect")
	}
	{
		result := IsNegativeZero(3.0)
		expect := false
		assert.Equal(t, result, expect, "doesn't match result & expect")
	}
	{
		result := IsNegativeZero(2.0)
		expect := false
		assert.Equal(t, result, expect, "doesn't match result & expect")
	}
	{
		result := IsNegativeZero(-5.0)
		expect := false
		assert.Equal(t, result, expect, "doesn't match result & expect")
	}
	{
		result := IsNegativeZero(math.Copysign(0, -1))
		expect := true
		assert.Equal(t, result, expect, "doesn't match result & expect")
	}

}

/*
  Expect(IsNegativeZero(n)).To(Equal(expected))
}

var _ = Describe("Tests", func() {
     It("Sample tests", func() {
       dotest(5.0, false)
       dotest(4.0, false)
       dotest(3.0, false)
       dotest(2.0, false)
       dotest(1.0, false)
       dotest(-5.0, false)
       dotest(-4.0, false)
       dotest(-2.0, false)
       dotest(-1.0, false)
       dotest(math.Inf(-1), false)
       dotest(math.Inf(1), false)
       dotest(0.0, false)
       dotest(math.NaN(), false)
       dotest(math.MaxFloat64, false)
       dotest(-math.MaxFloat64, false)
       dotest(math.SmallestNonzeroFloat64, false)
       dotest(math.Copysign(0, -1), true)
     })
})
*/
