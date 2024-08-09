package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSizeToNumber(t *testing.T) {
	{
		result, resBool := SizeToNumber("")
		expected, expBool := 0, false
		assert.Equal(t, result, expected, "Сравнение номеров")
		assert.Equal(t, resBool, expBool, "Ошибка логики")
	}
	{
		result, resBool := SizeToNumber("xm")
		expected, expBool := 0, false
		assert.Equal(t, result, expected, "Сравнение номеров")
		assert.Equal(t, resBool, expBool, "Ошибка логики")
	}
	{
		result, resBool := SizeToNumber("xxxm")
		expected, expBool := 0, false
		assert.Equal(t, result, expected, "Сравнение номеров")
		assert.Equal(t, resBool, expBool, "Ошибка логики")
	}
	{
		result, resBool := SizeToNumber("xxxx")
		expected, expBool := 0, false
		assert.Equal(t, result, expected, "Сравнение номеров")
		assert.Equal(t, resBool, expBool, "Ошибка логики")
	}
	{
		result, resBool := SizeToNumber("ssss")
		expected, expBool := 0, false
		assert.Equal(t, result, expected, "Сравнение номеров")
		assert.Equal(t, resBool, expBool, "Ошибка логики")
	}
	{
		result, resBool := SizeToNumber("hello world")
		expected, expBool := 0, false
		assert.Equal(t, result, expected, "Сравнение номеров")
		assert.Equal(t, resBool, expBool, "Ошибка логики")
	}
	result, resBool := SizeToNumber("xxxxxxxxxxxxxxxxxxxxxxxxxxxl")
	expected, expBool := 94, true
	assert.Equal(t, result, expected, "Сравнение номеров")
	assert.Equal(t, resBool, expBool, "Ошибка логики")
}

/*
dotest("", 0, false, "Blank string is invalid ('')")
        dotest("xm", 0, false, "Cannot apply modifier for medium size ('xm')")
        dotest("xxxm", 0, false, "Cannot apply modifier for medium size ('xxxm')")
        dotest("xxxx", 0, false, "No base size provided ('xxxx')")
        dotest("ssss", 0, false, "Only one base size is allowed ('ssss')")
        dotest("hello world", 0, false, "Any other text is invalid ('Hello world') is invalid")
        dotest("sm", 0, false, "'sm' should be invalid (two bases)")
        dotest("ml", 0, false, "'ml' should be invalid (two bases)")
        dotest("lm", 0, false, "'lm' should be invalid (two bases)")
        dotest("lx", 0, false, "'lx' should be invalid (base should be last)")
*/

/*
Random valid modifiers + base. Testing for 'xxxxxxxxxxxxxxxxxxxxxxxxxxxl'
Expected
    <int>: 0
to equal
    <int>: 94
*/
