package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDisemvowel(t *testing.T) {

	{
		result := Disemvowel("This website is for losers LOL!")
		expect := "Ths wbst s fr lsrs LL!"
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := Disemvowel("No offense but,\nYour writing is among the worst I've ever read")
		expect := "N ffns bt,\nYr wrtng s mng th wrst 'v vr rd"
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := Disemvowel("I couldn't BELIEVE it!")
		expect := " cldn't BLV t!"
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := Disemvowel("What are you, a communist?")
		expect := "Wht r y,  cmmnst?"
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
}
func TestBestPractice(t *testing.T) {
	{
		result := BestPractice("This website is for losers LOL!")
		expect := "Ths wbst s fr lsrs LL!"
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := BestPractice("No offense but,\nYour writing is among the worst I've ever read")
		expect := "N ffns bt,\nYr wrtng s mng th wrst 'v vr rd"
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := BestPractice("I couldn't BELIEVE it!")
		expect := " cldn't BLV t!"
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := BestPractice("What are you, a communist?")
		expect := "Wht r y,  cmmnst?"
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
}
func TestDisemvowelV3(t *testing.T) {
	{
		result := DisemvowelV3("This website is for losers LOL!")
		expect := "Ths wbst s fr lsrs LL!"
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := DisemvowelV3("No offense but,\nYour writing is among the worst I've ever read")
		expect := "N ffns bt,\nYr wrtng s mng th wrst 'v vr rd"
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := DisemvowelV3("I couldn't BELIEVE it!")
		expect := " cldn't BLV t!"
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := DisemvowelV3("What are you, a communist?")
		expect := "Wht r y,  cmmnst?"
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
}
