package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {

	{
		result := Solve("Codewars@codewars123.com")
		expect := []int{1, 18, 3, 2}
		assert.ElementsMatch(t, result, expect, "result value doesn't match expect value")
	}

	{
		result := Solve("bgA5<1d-tOwUZTS8yQ")
		expect := []int{7, 6, 3, 2}
		assert.ElementsMatch(t, result, expect, "result value doesn't match expect value")
	}

	{
		result := Solve("P*K4%>mQUDaG$h=cx2?.Czt7!Zn16p@5H")
		expect := []int{9, 9, 6, 9}
		assert.ElementsMatch(t, result, expect, "result value doesn't match expect value")
	}

	{
		result := Solve("RYT'>s&gO-.CM9AKeH?,5317tWGpS<*x2ukXZD")
		expect := []int{15, 8, 6, 9}
		assert.ElementsMatch(t, result, expect, "result value doesn't match expect value")
	}

	{
		result := Solve("$Cnl)Sr<7bBW-&qLHI!mY41ODe")
		expect := []int{10, 7, 3, 6}
		assert.ElementsMatch(t, result, expect, "result value doesn't match expect value")
	}
}

/*
  It("It should work for basic tests", func() {
        dotest("Codewars@codewars123.com", []int{1,18,3,2})
        dotest("bgA5<1d-tOwUZTS8yQ", []int{7,6,3,2})
        dotest("P*K4%>mQUDaG$h=cx2?.Czt7!Zn16p@5H", []int{9,9,6,9})
        dotest("RYT'>s&gO-.CM9AKeH?,5317tWGpS<*x2ukXZD", []int{15,8,6,9})
        dotest("$Cnl)Sr<7bBW-&qLHI!mY41ODe", []int{10,7,3,6})
  })
*/
