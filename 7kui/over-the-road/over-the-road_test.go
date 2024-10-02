package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOverTheRoad(t *testing.T) {

	{
		result := OverTheRoad(1, 3)
		expected := 6
		assert.Equal(t, result, expected, "value of result doesn't equal expected")
	}
	{
		result := OverTheRoad(3, 3)
		expected := 4
		assert.Equal(t, result, expected, "value of result doesn't equal expected")
	}
	{
		result := OverTheRoad(2, 3)
		expected := 5
		assert.Equal(t, result, expected, "value of result doesn't equal expected")
	}
	{
		result := OverTheRoad(3, 5)
		expected := 8
		assert.Equal(t, result, expected, "value of result doesn't equal expected")
	}
	{
		result := OverTheRoad(7, 11)
		expected := 16
		assert.Equal(t, result, expected, "value of result doesn't equal expected")
	}
	{
		result := OverTheRoad(23633656673, 310027696726)
		expected := 596421736780
		assert.Equal(t, result, expected, "value of result doesn't equal expected")
	}
}
