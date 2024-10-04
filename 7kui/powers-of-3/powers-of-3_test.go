package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLargestPower(t *testing.T) {
	{
		result := LargestPower(81)
		expect := 3
		assert.Equal(t, result, expect, "result value doesn't equal expect value")

	}
	{
		result := LargestPower(59049)
		expect := 9
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := LargestPower(59050)
		expect := 10
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := LargestPower(123456789)
		expect := 16
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := LargestPower(987654321)
		expect := 18
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}

}
func TestLargestPower1(t *testing.T) {
	{
		result := LargestPower1(81)
		expect := 3
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := LargestPower1(59049)
		expect := 9
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := LargestPower1(59050)
		expect := 10
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := LargestPower1(123456789)
		expect := 16
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := LargestPower1(987654321)
		expect := 18
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
}

/*
	Context("Small cases",func() {
		small_in := [...]uint64{2,3,4,5,7}
		small_ot := [...]int{0,0,1,1,1}
		for i,v := range small_in {run_test(v,small_ot[i])}})

	Context("Large cases",func() {
		large_in := [...]uint64{81,82,59049,59050,123456789,987654321}
		large_ot := [...]int{3,4,9,10,16,18}
*/
