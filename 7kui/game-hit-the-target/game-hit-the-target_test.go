package kata

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	//fmt.Println(Solution([][]rune{{'>', ' '}, {' ', 'x'}}))
	fmt.Println(Solution([][]rune{{' ', ' ', ' ', ' ', ' '},
		{' ', ' ', ' ', ' ', ' '},
		{' ', ' ', ' ', ' ', ' '},
		{' ', ' ', '>', ' ', 'x'},
		{' ', ' ', ' ', ' ', ' '}}))

	{
		result := Solution([][]rune{{'>', ' '}, {' ', 'x'}})
		expect := false
		assert.Equal(t, result, expect, "не соответвует result & expect")
	}
	{
		result := Solution([][]rune{{' ', ' ', ' ', ' ', ' '},
			{' ', ' ', ' ', ' ', ' '},
			{' ', ' ', ' ', ' ', ' '},
			{' ', ' ', '>', ' ', 'x'},
			{' ', ' ', ' ', ' ', ' '}})
		expect := true
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}

}

/*
       Expect(Solution([][]rune{{'>', ' '}, {' ', 'x'}})).To(Equal(false))
       Expect(Solution([][]rune{{' ', ' ', ' ', ' ', ' '},
  {' ', ' ', ' ', ' ', ' '},
  {' ', ' ', ' ', ' ', ' '},
  {' ', ' ', '>', ' ', 'x'},
  {' ', ' ', ' ', ' ', ' '}})).To(Equal(true))
     })
*/
