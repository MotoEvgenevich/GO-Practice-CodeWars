package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAngle(t *testing.T) {
	{
		result := Angle(3)
		var expect int = 180
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := Angle(4)
		var expect int = 360
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
}

/*
    It("Angle(3)", func() { Expect(Angle(3)).To(Equal(180)) })
    It("Angle(4)", func() { Expect(Angle(4)).To(Equal(360)) })
})
*/
