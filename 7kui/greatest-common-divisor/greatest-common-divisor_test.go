package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGcd(t *testing.T) {
	{
		result := Gcd(30, 12)
		var expect uint32 = 6
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := Gcd(8, 9)
		var expect uint32 = 1
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := Gcd(1, 1)
		var expect uint32 = 1
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
}

/*
TEST CASES:
dotest(30, 12, 6)
dotest(8, 9, 1)
dotest(1, 1, 1)
*/
