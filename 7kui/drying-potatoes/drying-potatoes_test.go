package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPotatoes(t *testing.T) {
	{
		result := Potatoes(99, 100, 98)
		expect := 50
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := Potatoes(82, 127, 80)
		expect := 114
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
}

/*
   dotest(99, 100, 98, 50)
   dotest(82, 127, 80, 114)
*/
