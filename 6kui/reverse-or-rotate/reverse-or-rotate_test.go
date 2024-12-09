package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRevrot(t *testing.T) {
	/* 	{
	   		result := Revrot("1234", 0)
	   		expected := ""
	   		assert.Equal(t, expected, result, "не соотвествуют данным:  dotest(\"1234\", 0, \"\")")
	   	}
	   	{
	   		result := Revrot("", 0)
	   		expected := ""
	   		assert.Equal(t, expected, result, "не соотвествуют данным:  dotest(\"\", 0, \"\")")
	   	}
	   	{
	   		result := Revrot("1234", 5)
	   		expected := ""
	   		assert.Equal(t, expected, result, "не соотвествуют данным:  dotest(\"1234\", 5, \"\")")
	   	} */
	{
		var s = "733049910872815764"
		result := Revrot(s, 5)
		expected := "330479108928157"
		assert.Equal(t, expected, result, "не соотвествуют данным:  dotest(s, 5, \"330479108928157\")")
	}

}

/*
   dotest("1234", 0, "")
   dotest("", 0, "")
   dotest("1234", 5, "")
   var s = "733049910872815764"
   dotest(s, 5, "330479108928157")
*/
