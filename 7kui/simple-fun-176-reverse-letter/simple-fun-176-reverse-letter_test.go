package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseLetters(t *testing.T) {
	{
		result := ReverseLetters("krishan")
		expected := "nahsirk"
		assert.Equal(t, result, expected, "данные не совпадают в тесте: krishan")
	}
	{
		result := ReverseLetters("ultr53o?n")
		expected := "nortlu"
		assert.Equal(t, result, expected, "данные не совпадают в тесте: ultr53o?n")
	}
	{
		result := ReverseLetters("ab23c")
		expected := "cba"
		assert.Equal(t, result, expected, "данные не совпадают в тесте: ab23c")
	}
	{
		result := ReverseLetters("krish21an")
		expected := "nahsirk"
		assert.Equal(t, result, expected, "данные не совпадают в тесте: krish21an")
	}

}

/*
   dotest("krishan","nahsirk")
   dotest("ultr53o?n","nortlu")
   dotest("ab23c","cba")
   dotest("krish21an","nahsirk")
   dotest("", "")
*/
