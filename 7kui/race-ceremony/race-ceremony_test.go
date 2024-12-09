package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXxx(t *testing.T) {
	{
		result := RacePodium(11)
		expect := [3]int{4, 5, 2}
		assert.Equal(t, expect, result, "result value doesn't equal expect value")
	}

	{
		result := RacePodium(6)
		expect := [3]int{2, 3, 1}
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := RacePodium(10)
		expect := [3]int{4, 5, 1}
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := RacePodium(100000)
		expect := [3]int{33334, 33335, 33331}
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := RacePodium(7)
		expect := [3]int{2, 4, 1}
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := RacePodium(8)
		expect := [3]int{3, 4, 1}
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}

}

/*
   dotest(11, [3]int{4, 5, 2})
   dotest(6, [3]int{2, 3, 1})
   dotest(10, [3]int{4, 5, 1})
   dotest(100000, [3]int{33334, 33335, 33331})
   dotest(7, [3]int{2, 4, 1})
   dotest(8, [3]int{3, 4, 1})
*/
