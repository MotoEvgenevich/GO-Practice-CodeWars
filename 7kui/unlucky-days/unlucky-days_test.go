package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnluckyDays(t *testing.T) {
	{
		result := UnluckyDays(2015)
		expect := 3
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := UnluckyDays(1986)
		expect := 1
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}

}
