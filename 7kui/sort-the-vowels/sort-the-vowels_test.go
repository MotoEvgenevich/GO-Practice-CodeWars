package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXxx(t *testing.T) {
	{
		result := SortVowels("Codewars")
		expect := "C|\n|o\nd|\n|e\nw|\n|a\nr|\ns|"
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := SortVowels("Rnd Te5T")
		expect := "R|\nn|\nd|\n |\nT|\n|e\n5|\nT|"
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := SortVowels("yo!")
		expect := "y|\n|o\n!|"
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := SortVowels("1337")
		expect := "1|\n3|\n3|\n7|"
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := SortVowels("")
		expect := ""
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := SortVowels("[)%")
		expect := "[|\n)|\n%|"
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := SortVowels("LOrEm IpsUm dOlOr sIt AmEt")
		expect := "L|\n|O\nr|\n|E\nm|\n |\n|I\np|\ns|\n|U\nm|\n |\nd|\n|O\nl|\n|O\nr|\n |\ns|\n|I\nt|\n |\n|A\nm|\n|E\nt|"
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
}

/*
	dotest("Codewars", "C|\n|o\nd|\n|e\nw|\n|a\nr|\ns|")
	dotest("Rnd Te5T", "R|\nn|\nd|\n |\nT|\n|e\n5|\nT|")
	dotest("yo!", "y|\n|o\n!|")
	dotest("1337", "1|\n3|\n3|\n7|")
	dotest("", "")
	dotest("[)%", "[|\n)|\n%|")
	dotest("LOrEm IpsUm dOlOr sIt AmEt", "L|\n|O\nr|\n|E\nm|\n |\n|I\np|\ns|\n|U\nm|\n |\nd|\n|O\nl|\n|O\nr|\n |\ns|\n|I\nt|\n |\n|A\nm|\n|E\nt|")
*/
