package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrder(t *testing.T) {
	{
		result := Order("is2 Thi1s T4est 3a")
		expect := "Thi1s is2 3a T4est"
		assert.Equal(t, expect, result, "result doesn't equal expect")
	}
	{
		result := Order("4of Fo1r pe6ople g3ood th5e the2")
		expect := "Fo1r the2 g3ood 4of th5e pe6ople"
		assert.Equal(t, expect, result, "result doesn't equal expect")
	}
	{
		result := Order("")
		expect := ""
		assert.Equal(t, expect, result, "result doesn't equal expect")
	}
}

/*
   Expect(Order("is2 Thi1s T4est 3a")).To(Equal("Thi1s is2 3a T4est"))
   Expect(Order("4of Fo1r pe6ople g3ood th5e the2")).To(Equal("Fo1r the2 g3ood 4of th5e pe6ople"))
   Expect(Order("")).To(Equal(""))
*/
