package kata

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHighAndLow(t *testing.T) {
	{
		result := HighAndLow("1 2 3 4 5")
		expect := "5 1"
		assert.Equal(t, result, expect, "не соответсвует значение result & expect")
	}
	{
		result := HighAndLow("1 2 -3 4 5")
		expect := "5 -3"
		fmt.Println(result)
		assert.Equal(t, result, expect, "не соответсвует значение result & expect")
	}
	{
		result := HighAndLow("1 9 3 4 -5")
		expect := "9 -5"
		fmt.Println(result)
		assert.Equal(t, result, expect, "не соответсвует значение result & expect")
	}

}

/*
HighAndLow("1 2 3 4 5")  // return "5 1"
HighAndLow("1 2 -3 4 5") // return "5 -3"
HighAndLow("1 9 3 4 -5") // return "9 -5"
*/
