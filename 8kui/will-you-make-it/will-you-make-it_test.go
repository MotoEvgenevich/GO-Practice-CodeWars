package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestZeroFuel(t *testing.T) {
	{
		result := ZeroFuel(50, 25, 2)
		expect := true
		assert.Equal(t, result, expect, "result value doesn't equal expect value")

	}
	{
		result := ZeroFuel(60, 30, 3)
		expect := true
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := ZeroFuel(70, 25, 1)
		expect := false
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := ZeroFuel(100, 25, 3)
		expect := false
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
}

/*
   DoTest(50, 25, 2, true)
   DoTest(60, 30, 3, true)
   DoTest(70, 25, 1, false)
   DoTest(100, 25, 3, false)
*/
