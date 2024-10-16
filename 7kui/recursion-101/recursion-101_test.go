package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRecursion(t *testing.T) {
	{
		result := Solve(6, 19)
		expect := []int{6, 7}
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}

	{
		result := Solve(2, 1)
		expect := []int{0, 1}
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}

	{
		result := Solve(22, 5)
		expect := []int{0, 1}
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}

	{
		result := Solve(2, 10)
		expect := []int{2, 2}
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}

	{
		result := Solve(8796203, 7556)
		expect := []int{1019, 1442}
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}

	{
		result := Solve(19394, 19394)
		expect := []int{19394, 19394}
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}

}

/*
   dotest(6,19, []int{6,7})
   dotest(2,1, []int{0,1})
   dotest(22,5, []int{0,1})
   dotest(2,10, []int{2,2})
   dotest(8796203,7556, []int{1019,1442})
   dotest(19394,19394, []int{19394,19394})
*/
