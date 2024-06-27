package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortMyString(t *testing.T) {

	{
		result := SortMyString("CodeWars")
		expect := "CdWr oeas"
		assert.Equal(t, result, expect, "не соотвествует")
	}
	{
		result := SortMyString("YCOLUE'VREER")
		expect := "YOU'RE CLEVER"
		assert.Equal(t, result, expect, "не соотвествует")
	}
}
