package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortNumbers(t *testing.T) {

	{
		result := SortNumbers([]int{1, 2, 10, 50, 5})
		expect := []int{1, 2, 5, 10, 50}
		assert.ElementsMatch(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := SortNumbers([]int{})
		expect := []int{}
		assert.ElementsMatch(t, result, expect, "result value doesn't equal expect value")
	}
}
func TestSortNumbersV2(t *testing.T) {
	{
		result := SortNumbersV2([]int{1, 2, 10, 50, 5})
		expect := []int{1, 2, 5, 10, 50}
		assert.ElementsMatch(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := SortNumbersV2([]int{})
		expect := []int{}
		assert.ElementsMatch(t, result, expect, "result value doesn't equal expect value")
	}
}

/*
   Expect(Expect(SortNumbers([]int{1, 2, 10, 50, 5})).To(Equal([]int{1,2,5,10,50})))
   Expect(Expect(SortNumbers([]int{})).To(Equal([]int{})))
*/
