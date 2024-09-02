package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalc(t *testing.T) {
	{
		result := Calc("aaaaaddddr")
		expect := 30
		assert.Equal(t, result, expect, "не соотвествует звачения result и expect")
	}
	{
		result := Calc("abcdef")
		expect := 6
		assert.Equal(t, result, expect, "не соотвествует звачения result и expect")
	}
	{
		result := Calc("ifkhchlhfd")
		expect := 6
		assert.Equal(t, result, expect, "не соотвествует звачения result и expect")
	}
	{
		result := Calc("jfmgklf8hglbe")
		expect := 6
		assert.Equal(t, result, expect, "не соотвествует звачения result и expect")
	}
	{
		result := Calc("jaam")
		expect := 12
		assert.Equal(t, result, expect, "не соотвествует звачения result и expect")
	}
}

/*

var _ = Describe("Tests", func() {
     It("Sample tests", func() {
       dotest("abcdef", 6)
       dotest("ifkhchlhfd", 6)
       dotest("aaaaaddddr", 30)
       dotest("jfmgklf8hglbe", 6)
       dotest("jaam", 12)
     })
})
*/
