package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuietTime(t *testing.T) {

	{
		result := QueueTime([]int{}, 1)
		expect := 0
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := QueueTime([]int{1, 2, 3, 4}, 1)
		expect := 10
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}

	{
		result := QueueTime([]int{2, 2, 3, 3, 4, 4}, 2)
		expect := 9
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}

	{
		result := QueueTime([]int{1, 2, 3, 4, 5}, 100)
		expect := 5
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}

}

/*
   Expect(QueueTime([]int{}, 1)).To(Equal(0))
   Expect(QueueTime([]int{1,2,3,4}, 1)).To(Equal(10))
   Expect(QueueTime([]int{2,2,3,3,4,4}, 2)).To(Equal(9))
   Expect(QueueTime([]int{1,2,3,4,5}, 100)).To(Equal(5))
*/
