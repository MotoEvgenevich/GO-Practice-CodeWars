package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreatePhoneNumber(t *testing.T) {
	{
		result := CreatePhoneNumber([10]uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 0})
		expect := "(123) 456-7890"
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
}

/*
   Expect(CreatePhoneNumber([10]uint{1,2,3,4,5,6,7,8,9,0})).To(Equal("(123) 456-7890"))
*/
