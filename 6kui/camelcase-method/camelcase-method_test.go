package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCamelCase(t *testing.T) {
	{
		actual := CamelCase("test case")
		expect := "TestCase"
		assert.Equal(t, expect, actual, "result doesn't equal expect")
	}
}

/*
   {"test case","TestCase"},
   {"camel case method","CamelCaseMethod"},
   {"say hello ","SayHello"},
   {" camel case word","CamelCaseWord"},
   {"",""},
*/
