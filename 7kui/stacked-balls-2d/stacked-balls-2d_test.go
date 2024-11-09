package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStackHeight2d(t *testing.T) {
	{
		result := StackHeight2d(1)
		expect := 1.0
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := StackHeight2d(2)
		expect := 1.866
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := StackHeight2d(0)
		expect := 0.0
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
}

/*
   It("test0", func() {
       DoTest(0, 0.0)
   })
   It("test1", func() {
       DoTest(1, 1.0)
   })
   It("test2", func() {
       DoTest(2, 1.866)
*/
