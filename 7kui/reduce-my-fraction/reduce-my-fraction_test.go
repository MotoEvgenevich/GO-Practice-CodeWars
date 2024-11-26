package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReduceFraction(t *testing.T) {
	{
		result := ReduceFraction([2]int{60, 20})
		expect := [2]int{3, 1}
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := ReduceFraction([2]int{80, 120})
		expect := [2]int{2, 3}
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := ReduceFraction([2]int{4, 2})
		expect := [2]int{2, 1}
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := ReduceFraction([2]int{45, 120})
		expect := [2]int{3, 8}
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := ReduceFraction([2]int{1000, 1})
		expect := [2]int{1000, 1}
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := ReduceFraction([2]int{1, 1})
		expect := [2]int{1, 1}
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
}

/*
   dotest([2]int{60, 20}, [2]int{3, 1});
   dotest([2]int{80, 120}, [2]int{2, 3});
   dotest([2]int{4, 2}, [2]int{2, 1});
   dotest([2]int{45, 120}, [2]int{3, 8});
   dotest([2]int{1000, 1}, [2]int{1000, 1});
   dotest([2]int{1, 1}, [2]int{1, 1});
*/
