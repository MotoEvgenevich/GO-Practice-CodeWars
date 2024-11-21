package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNbDig(t *testing.T) {
	{
		result := NbDig(10, 1)
		expect := 4
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}

	{
		result := NbDig(5750, 0)
		expect := 4700
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := NbDig(550, 5)
		expect := 213
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}

}

/*
   dotest(550, 5, 213)
   dotest(5750, 0, 4700)
*/
