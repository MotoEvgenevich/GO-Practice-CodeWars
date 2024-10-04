package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountRedBeads(t *testing.T) {
	{
		result := CountRedBeads(0)
		var expect int = 0
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := CountRedBeads(1)
		var expect int = 0
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := CountRedBeads(3)
		var expect int = 4
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := CountRedBeads(5)
		var expect int = 8
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
}

/*
    Expect(CountRedBeads(0)).To(Equal(0))
    Expect(CountRedBeads(1)).To(Equal(0))
    Expect(CountRedBeads(3)).To(Equal(4))
    Expect(CountRedBeads(5)).To(Equal(8))
  })
*/
