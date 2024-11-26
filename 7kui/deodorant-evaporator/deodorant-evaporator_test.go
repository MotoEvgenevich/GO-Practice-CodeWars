package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEvaporator(t *testing.T) {
	{
		result := Evaporator(10, 10, 10)
		expect := 22
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := Evaporator(10, 10, 5)
		expect := 29
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := Evaporator(100, 5, 5)
		expect := 59
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}

}

/*
   dotest(10, 10, 10, 22)
   dotest(10, 10, 5, 29)
   dotest(100, 5, 5, 59)
*/
