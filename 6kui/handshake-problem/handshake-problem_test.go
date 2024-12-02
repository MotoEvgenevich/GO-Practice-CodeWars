package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXxx(t *testing.T) {
	{
		result := GetParticipants(0)
		expect := 0
		assert.Equal(t, expect, result, "result doesn't equal expect")
	}
	{
		result := GetParticipants(1)
		expect := 2
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := GetParticipants(3)
		expect := 3
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := GetParticipants(6)
		expect := 4
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := GetParticipants(7)
		expect := 5
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}

}

/*
   It("should pass for GetParticipants(0)", func() { Expect(GetParticipants(0)).To(Equal(0)) })
   It("should pass for GetParticipants(1)", func() { Expect(GetParticipants(1)).To(Equal(2)) })
   It("should pass for GetParticipants(3)", func() { Expect(GetParticipants(3)).To(Equal(3)) })
   It("should pass for GetParticipants(6)", func() { Expect(GetParticipants(6)).To(Equal(4)) })
   It("should pass for GetParticipants(7)", func() { Expect(GetParticipants(7)).To(Equal(5)) })
*/
