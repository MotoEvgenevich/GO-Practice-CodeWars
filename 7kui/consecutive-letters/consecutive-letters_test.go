package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCoprime(t *testing.T) {
	{
		result := Solve("abc")
		expect := true
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := Solve("abd")
		expect := false
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := Solve("dabc")
		expect := true
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := Solve("abbc")
		expect := false
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}

}

func TestSolve(t *testing.T) {
	cases := []struct {
		input string
		want  bool
	}{
		{"abc", true},
		{"abd", false},
		{"dabc", true},
		{"abbc", false},
		{"v", true},
	}

	for _, tc := range cases {
		got := Solve(tc.input)
		if got != tc.want {
			t.Errorf("Solve(%q) = %v, want %v", tc.input, got, tc.want)
		}
	}
}

/*

   dotest("abc", true)
   dotest("abd", false)
   dotest("dabc", true)
   dotest("abbc", false)

*/
