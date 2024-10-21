package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhatCentury(t *testing.T) {
	{
		expected := "20th"
		result := WhatCentury("1999")
		assert.Equal(t, expected, result, "не соотвествуют данным:  dotest(\"1999\", \"20th\")")
	}
	{
		expected := "21st"
		result := WhatCentury("2011")
		assert.Equal(t, expected, result, "не соотвествуют данным:  dotest(\"2011\", \"21st\")")
	}
	{
		expected := "22nd"
		result := WhatCentury("2154")
		assert.Equal(t, expected, result, "не соотвествуют данным:  dotest(\"2154\", \"22nd\")")
	}
	{
		expected := "23rd"
		result := WhatCentury("2259")
		assert.Equal(t, expected, result, "не соотвествуют данным:  dotest(\"2259\", \"23rd\")")
	}
	{
		expected := "13th"
		result := WhatCentury("1234")
		assert.Equal(t, expected, result, "не соотвествуют данным:  dotest(\"1234\", \"12th\")")
	}
	{
		expected := "11th"
		result := WhatCentury("1023")
		assert.Equal(t, expected, result, "не соотвествуют данным:  dotest(\"1023\", \"11th\")")
	}
}

/*
   dotest("2011", "21st")
   dotest("2154", "22nd")
   dotest("2259", "23rd")
   dotest("1234", "13th")
   dotest("1023", "11th")
*/
