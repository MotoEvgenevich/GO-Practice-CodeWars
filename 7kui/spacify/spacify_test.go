package kata

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSpacify(t *testing.T) {
	{
		result := Spacify("hello world")
		expect := "h e l l o   w o r l d"
		assert.Equal(t, result, expect, "doesn't match result & expect")
		fmt.Println(result)
	}
	{
		result := Spacify("12345")
		expect := "1 2 3 4 5"
		assert.Equal(t, result, expect, "doesn't match result & expect")
		fmt.Println(result)
	}
	{
		result := Spacify("Pippi")
		expect := "P i p p i"
		assert.Equal(t, result, expect, "doesn't match result & expect")
		fmt.Println(result)
	}
	{
		result := Spacify("a")
		expect := "a"
		assert.Equal(t, result, expect, "doesn't match result & expect")
		fmt.Println(result)
	}
	{
		result := Spacify("")
		expect := ""
		assert.Equal(t, result, expect, "doesn't match result & expect")
		fmt.Println(result)
	}
}

/*
   It("Sample tests", func() {
      dotest("hello world", "h e l l o   w o r l d")
      dotest("12345", "1 2 3 4 5")
      dotest("Pippi", "P i p p i")
      dotest("a", "a")
      dotest("", "")
   })
*/
