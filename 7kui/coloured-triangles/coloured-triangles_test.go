package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTriangle(t *testing.T) {
	{
		result := Triangle("GB")
		expect := 'R'
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := Triangle("R")
		expect := 'R'
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := Triangle("RGBG")
		expect := 'B'
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := Triangle("RBRGBRB")
		expect := 'G'
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := Triangle("RBRGBRBGGRRRBGBBBGG")
		expect := 'G'
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
}

/*
   dotest("GB", 'R')
   dotest("RRR", 'R')
   dotest("RGBG", 'B')
   dotest("RBRGBRB", 'G')
   dotest("RBRGBRBGGRRRBGBBBGG", 'G')
   dotest("B", 'B')
*/
