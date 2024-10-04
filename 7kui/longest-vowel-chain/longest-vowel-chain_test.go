package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	{
		result := Solve("codewarriors")
		expect := 2
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := Solve("suoidea")
		expect := 3
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := Solve("ultrarevolutionariees")
		expect := 3
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := Solve("strengthlessnesses")
		expect := 1
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := Solve("cuboideonavicuare")
		expect := 2
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := Solve("chrononhotonthuooaos")
		expect := 5
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := Solve("iiihoovaeaaaoougjyaw")
		expect := 8
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
}

/*
TESTS:
		 Expect(Solve("codewarriors")).To(Equal(2))
        Expect(Solve("suoidea")).To(Equal(3))
        Expect(Solve("ultrarevolutionariees")).To(Equal(3))
        Expect(Solve("strengthlessnesses")).To(Equal(1))
        Expect(Solve("cuboideonavicuare")).To(Equal(2))
        Expect(Solve("chrononhotonthuooaos")).To(Equal(5))
        Expect(Solve("iiihoovaeaaaoougjyaw")).To(Equal(8))

*/
