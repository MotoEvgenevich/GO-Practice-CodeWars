package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	{
		actual := Solve("your code rocks")
		expect := "skco redo cruoy"
		assert.Equal(t, expect, actual, "doesnt match data")
	}
	{
		actual := Solve("codewars")
		expect := "srawedoc"
		assert.Equal(t, expect, actual, "doesnt match data")
	}
	{
		actual := Solve("your code")
		expect := "edoc ruoy"
		assert.Equal(t, expect, actual, "doesnt match data")
	}
}

/*
   dotest("codewars","srawedoc")
   dotest("your code","edoc ruoy")
   dotest("your code rocks","skco redo cruoy")
*/
