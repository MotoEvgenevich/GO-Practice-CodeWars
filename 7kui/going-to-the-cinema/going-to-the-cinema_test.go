package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMovie(t *testing.T) {
	{
		result := Movie(500, 15, 0.9)
		expect := 43
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := Movie(0, 10, 0.95)
		expect := 2
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
}

/*
   dotest(500, 15, 0.9, 43)
   dotest(0, 10, 0.95, 2)
*/
