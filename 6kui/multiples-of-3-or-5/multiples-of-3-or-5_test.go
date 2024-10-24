package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiple3And5(t *testing.T) {
	{
		result := Multiple3And5(10)
		expect := 23
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}

}

/*
   Expect(Multiple3And5(10)).To(Equal(23))
*/
