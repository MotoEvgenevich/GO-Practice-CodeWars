package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {

	{
		result := Solve("123", 1)
		expect := 23
		assert.Equal(t, expect, result, "result value doesn't equal expect value")
	}

	{
		result := Solve("1234", 1)
		expect := 234
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}

	{
		result := Solve("1234", 2)
		expect := 34
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}

	{
		result := Solve("1234", 3)
		expect := 4
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}

}

/*
   dotest("123", 1, 23)
   dotest("1234",1, 234)
   dotest("1234",2, 34)
   dotest("1234",3, 4)
*/
