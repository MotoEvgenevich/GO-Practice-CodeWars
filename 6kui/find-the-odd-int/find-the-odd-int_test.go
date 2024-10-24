package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindOdd(t *testing.T) {
	{
		result := FindOdd([]int{20, 1, -1, 2, -2, 3, 3, 5, 5, 1, 2, 4, 20, 4, -1, -2, 5})
		expect := 5
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}

	{
		result := FindOdd([]int{1, 1, 2, -2, 5, 2, 4, 4, -1, -2, 5})
		expect := -1
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}

	{
		result := FindOdd([]int{20, 1, 1, 2, 2, 3, 3, 5, 5, 4, 20, 4, 5})
		expect := 5
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}

	{
		result := FindOdd([]int{10})
		expect := 10
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}

	{
		result := FindOdd([]int{1, 1, 1, 1, 1, 1, 10, 1, 1, 1, 1})
		expect := 10
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}

	{
		result := FindOdd([]int{5, 4, 3, 2, 1, 5, 4, 3, 2, 10, 10})
		expect := 1
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}

}

/*

   arr := [...][]int{
       {20,1,-1,2,-2,3,3,5,5,1,2,4,20,4,-1,-2,5},
       {1,1,2,-2,5,2,4,4,-1,-2,5},
       {20,1,1,2,2,3,3,5,5,4,20,4,5},
       {10},
       {1,1,1,1,1,1,10,1,1,1,1},
       {5,4,3,2,1,5,4,3,2,10,10},
   }
   sol := [...]int{5,-1,5,10,10,1}
*/
