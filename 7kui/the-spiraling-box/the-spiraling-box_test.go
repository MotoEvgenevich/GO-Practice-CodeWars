package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateBox(t *testing.T) {
	{
		result := CreateBox(7, 8)
		expect := [][]int{{1, 1, 1, 1, 1, 1, 1}, {1, 2, 2, 2, 2, 2, 1}, {1, 2, 3, 3, 3, 2, 1}, {1, 2, 3, 4, 3, 2, 1}, {1, 2, 3, 4, 3, 2, 1}, {1, 2, 3, 3, 3, 2, 1}, {1, 2, 2, 2, 2, 2, 1}, {1, 1, 1, 1, 1, 1, 1}}
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := CreateBox(8, 7)
		expect := [][]int{{1, 1, 1, 1, 1, 1, 1, 1}, {1, 2, 2, 2, 2, 2, 2, 1}, {1, 2, 3, 3, 3, 3, 2, 1}, {1, 2, 3, 4, 4, 3, 2, 1}, {1, 2, 3, 3, 3, 3, 2, 1}, {1, 2, 2, 2, 2, 2, 2, 1}, {1, 1, 1, 1, 1, 1, 1, 1}}
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
}

/*
    Expect(CreateBox(7, 8)).To(Equal(box_7_8))
    Expect(CreateBox(8, 7)).To(Equal(box_8_7))
    Expect(CreateBox(4, 2)).To(Equal(box_4_2))
    Expect(CreateBox(2, 4)).To(Equal(box_2_4))
  })
*/
