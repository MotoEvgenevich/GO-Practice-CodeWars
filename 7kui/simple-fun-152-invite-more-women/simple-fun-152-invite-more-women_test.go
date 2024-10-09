package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInviteMoreWomen(t *testing.T) {

	{
		result := inviteMoreWomen([]int{1, -1, 1})
		expect := true
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := inviteMoreWomen([]int{1, 1, 1})
		expect := true
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := inviteMoreWomen([]int{-1, -1, -1})
		expect := false
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := inviteMoreWomen([]int{1, -1})
		expect := false
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
}

/*
	Expect(inviteMoreWomen([]int{1, -1, 1})).To(Equal(true))
	Expect(inviteMoreWomen([]int{1, 1, 1})).To(Equal(true))
	Expect(inviteMoreWomen([]int{-1, -1, -1})).To(Equal(false))
	Expect(inviteMoreWomen([]int{1, -1})).To(Equal(false))
*/
