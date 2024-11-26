package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeclareWinner(t *testing.T) {
	{
		result := DeclareWinner(Fighter{"Lew", 10, 2}, Fighter{"Harry", 5, 4}, "Lew")
		expect := "Lew"
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := DeclareWinner(Fighter{"Lew", 10, 2}, Fighter{"Harry", 5, 4}, "Harry")
		expect := "Harry"
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := DeclareWinner(Fighter{"Harald", 20, 5}, Fighter{"Harry", 5, 4}, "Harry")
		expect := "Harald"
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := DeclareWinner(Fighter{"Harald", 20, 5}, Fighter{"Harry", 5, 4}, "Harald")
		expect := "Harald"
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := DeclareWinner(Fighter{"Jerry", 30, 3}, Fighter{"Harald", 20, 5}, "Jerry")
		expect := "Harald"
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := DeclareWinner(Fighter{"Jerry", 30, 3}, Fighter{"Harald", 20, 5}, "Harald")
		expect := "Harald"
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
}

/*
	Expect(DeclareWinner(Fighter{"Lew", 10, 2}, Fighter{"Harry", 5, 4}, "Lew")).To(Equal("Lew"))
	Expect(DeclareWinner(Fighter{"Lew", 10, 2}, Fighter{"Harry", 5, 4}, "Harry")).To(Equal("Harry"))
	Expect(DeclareWinner(Fighter{"Harald", 20, 5}, Fighter{"Harry", 5, 4}, "Harry")).To(Equal("Harald"))
	Expect(DeclareWinner(Fighter{"Harald", 20, 5}, Fighter{"Harry", 5, 4}, "Harald")).To(Equal("Harald"))
	Expect(DeclareWinner(Fighter{"Jerry", 30, 3}, Fighter{"Harald", 20, 5}, "Jerry")).To(Equal("Harald"))
	Expect(DeclareWinner(Fighter{"Jerry", 30, 3}, Fighter{"Harald", 20, 5}, "Harald")).To(Equal("Harald"))
*/
