package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSeven(t *testing.T) {

	{
		result := Seven(477557101)
		expect := []int{28, 7}
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}

	{
		result := Seven(477557102)
		expect := []int{47, 7}
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
}

/*
    It("should handle basic cases", func() {
        dotest(477557101, []int{28, 7})
        dotest(477557102, []int{47, 7})

    })
})
*/
