package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLateRide(t *testing.T) {
	{
		result := LateRide(240)
		expect := 4
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := LateRide(808)
		expect := 14
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := LateRide(1439)
		expect := 19
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := LateRide(0)
		expect := 0
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := LateRide(23)
		expect := 5
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := LateRide(8)
		expect := 8
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}

}

/*
   Testing(240, 4)
   Testing(808, 14)
   Testing(1439, 19)
   Testing(0, 0)
   Testing(23, 5)
   Testing(8, 8)
*/
