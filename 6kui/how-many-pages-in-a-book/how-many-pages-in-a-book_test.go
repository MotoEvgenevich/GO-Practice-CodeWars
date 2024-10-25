package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAmountOfPages(t *testing.T) {
	{
		result := AmountOfPages(5)
		expected := 5
		assert.Equal(t, expected, result, "не соотвествуют данным:  dotest(5, 5)")
	}
	{
		result := AmountOfPages(25)
		expected := 17
		assert.Equal(t, expected, result, "не соотвествуют данным:  dotest(25, 17)")
	}
	{
		result := AmountOfPages(1095)
		expected := 401
		assert.Equal(t, expected, result, "не соотвествуют данным:  dotest(1095, 401)")
	}
	{
		result := AmountOfPages(185)
		expected := 97
		assert.Equal(t, expected, result, "не соотвествуют данным:  dotest(185, 97)")
	}
	{
		result := AmountOfPages(660)
		expected := 256
		assert.Equal(t, expected, result, "не соотвествуют данным:  dotest(660, 256)")
	}

	{
		result := AmountOfPages(33067)
		expected := 8543
		assert.Equal(t, expected, result, "не соотвествуют данным:  dotest(0, 0)")
	}
}

/*
   dotest(5, 5);
   dotest(25, 17)
   dotest(1095, 401)
   dotest(185, 97)
   dotest(660, 256)
   With summary = 33067
Expected
    <int>: 8544
to equal
    <int>: 8543

*/
