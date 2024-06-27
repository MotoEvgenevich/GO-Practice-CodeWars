package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthOfSequence(t *testing.T) {

	{
		result := LengthOfSequence([]int{0, -3, 7, 4, 0, 3, 7, 9}, 7)
		expect := 5
		assert.Equal(t, result, expect, "не соотвествует result & expect")

	}
	{
		result := LengthOfSequence([]int{}, 1)
		expect := 0
		assert.Equal(t, result, expect, "не соотвествует result & expect")
	}

	{
		result := LengthOfSequence([]int{86, -93, -43, 59, 85, 10, -38, 98, 19, -58, 86, -65, -49, -57, -12, -73, 45}, 86)
		expect := 11
		assert.Equal(t, result, expect, "не соотвествует result & expect")
	}
	{
		result := LengthOfSequence([]int{7, 1, 7, 1, 7}, 7)
		expect := 0
		assert.Equal(t, result, expect, "не соотвествует result & expect")
	}
	{
		result := LengthOfSequence([]int{0, -3, 7, 4, 4, 3, 7, 9}, 4)
		expect := 2
		assert.Equal(t, result, expect, "не соотвествует result & expect")
	}

}

/*
  It("No occurrence of key", func() {
    dotest([]int{}, 1, 0)
    dotest([]int{1, 2, 3}, 4, 0)
  })

  It("One occurrence of key", func() {
    dotest([]int{1}, 1, 0)
    dotest([]int{1, 2, 3, 4, 5}, 3, 0)
  })

  It("Two occurrences of key", func() {
    dotest([]int{0, -3, 7, 4, 0, 3, 7, 9}, 7, 5)
  })

  It("Three occurrences of key", func() {
    dotest([]int{7, 1, 7, 1, 7}, 7, 0)
  })
})

CRASHED TESTS:
With arr=[9 -3 7 4 0 3 7 9], key=9
Expected
    <int>: 0
to equal
    <int>: 8

With arr=[86 -93 -43 59 85 10 -38 98 19 -58 86 -65 -49 -57 -12 -73 45], key=86
Expected
    <int>: 0
to equal
    <int>: 11

With arr=[7 1 7 1 7], key=7
Expected
    <int>: 3
to equal
    <int>: 0


    With arr=[0 -3 7 4 4 3 7 9], key=4
Expected
    <int>: 0
to equal
    <int>: 2
*/
