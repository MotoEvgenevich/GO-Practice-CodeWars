package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXxx(t *testing.T) {
	{
		result := Vaporcode("Lets go to the movies")
		expect := "L  E  T  S  G  O  T  O  T  H  E  M  O  V  I  E  S"
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := Vaporcode("Why isnt my code working")
		expect := "W  H  Y  I  S  N  T  M  Y  C  O  D  E  W  O  R  K  I  N  G"
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
}

/*
   dotest("Lets go to the movies", "L  E  T  S  G  O  T  O  T  H  E  M  O  V  I  E  S")
   dotest("Why isnt my code working", "W  H  Y  I  S  N  T  M  Y  C  O  D  E  W  O  R  K  I  N  G")
*/
