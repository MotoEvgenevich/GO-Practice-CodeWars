package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBestTravel(t *testing.T) {
	{
		actual := ChooseBestSum(163, 3, []int{50, 55, 56, 57, 58})
		expect := 163
		assert.Equal(t, expect, actual, "result doesn't equal expect")
	}
	{
		actual := ChooseBestSum(163, 3, []int{50})
		expect := -1
		assert.Equal(t, expect, actual, "result doesn't equal expect")
	}
	{
		actual := ChooseBestSum(230, 3, []int{91, 74, 73, 85, 73, 81, 87})
		expect := 228
		assert.Equal(t, expect, actual, "result doesn't equal expect")
	}
	{
		actual := ChooseBestSum(331, 2, []int{91, 74, 73, 85, 73, 81, 87})
		expect := 178
		assert.Equal(t, expect, actual, "result doesn't equal expect")
	}
	{
		actual := ChooseBestSum(331, 4, []int{91, 74, 73, 85, 73, 81, 87})
		expect := 331
		assert.Equal(t, expect, actual, "result doesn't equal expect")
	}
	{
		actual := ChooseBestSum(331, 5, []int{91, 74, 73, 85, 73, 81, 87})
		expect := -1
		assert.Equal(t, expect, actual, "result doesn't equal expect")
	}
}

/*
   var ts = []int{50, 55, 56, 57, 58}
   dotest(163, 3, ts, 163)
   ts = []int{50}
   dotest(163, 3, ts, -1)
   ts = []int{91, 74, 73, 85, 73, 81, 87}
   dotest(230, 3, ts, 228)
   dotest(331, 2, ts, 178)
   dotest(331, 4, ts, 331)
   dotest(331, 5, ts, -1)
*/
