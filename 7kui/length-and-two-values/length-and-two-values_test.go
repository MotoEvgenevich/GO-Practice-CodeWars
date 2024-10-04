package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAlternate(t *testing.T) {

	{
		result := Alternate(5, "true", "false")
		expect := []string{"true", "false", "true", "false", "true"}
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := Alternate(10, "blue", "red")
		expect := []string{"blue", "red", "blue", "red", "blue", "red", "blue", "red", "blue", "red"}
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := Alternate(0, "one", "two")
		expect := []string{}
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := Alternate(20, "blue", "red")
		expect := []string{"blue", "red", "blue", "red", "blue", "red", "blue", "red", "blue", "red", "blue", "red", "blue", "red", "blue", "red", "blue", "red", "blue", "red"}
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
}

/*
     Expect(Alternate(5, "true", "false")).To(Equal([]string{"true", "false", "true", "false", "true"}))
     Expect(Alternate(20, "blue", "red")).To(Equal([]string{"blue", "red", "blue", "red", "blue", "red", "blue", "red", "blue", "red", "blue", "red", "blue", "red", "blue", "red", "blue", "red", "blue", "red"}))
     Expect(Alternate(0, "", "")).To(BeEmpty())
   })
*/
