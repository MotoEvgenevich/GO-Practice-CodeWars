package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncodeWithMap(t *testing.T) {
	{
		str := "Ala has a cat"
		result := EncodeWithMap(str)
		assert.Equal(t, result, "Gug hgs g cgt", "Ожидаемая строка не равна результату")
	}
	{
		str := "ABCD"
		result := EncodeWithMap(str)
		assert.Equal(t, result, "GBCE", "Ожидаемая строка не равна результату")
	}
	{
		str := "gaderypoluki"
		result := EncodeWithMap(str)
		assert.Equal(t, result, "agedyropulik", "Ожидаемая строка не равна результату")
	}
	{
		str := "Gug hgs g cgt"
		result := EncodeWithMap(str)
		assert.Equal(t, result, "Ala has a cat", "Ожидаемая строка не равна результату")
	}
	{
		str := "GBCE"
		result := EncodeWithMap(str)
		assert.Equal(t, result, "ABCD", "Ожидаемая строка не равна результату")
	}
	{
		str := "agedyropulik"
		result := EncodeWithMap(str)
		assert.Equal(t, result, "gaderypoluki", "Ожидаемая строка не равна результату")
	}

}

/* var _ = Describe("Tests", func() {
	It("Basic tests", func() {
		dotestEncode("Ala has a cat", "Gug hgs g cgt")
		dotestEncode("ABCD", "GBCE")
		dotestEncode("gaderypoluki", "agedyropulik")

		dotestDecode("Gug hgs g cgt", "Ala has a cat")
		dotestDecode("GBCE", "ABCD")
		dotestDecode("agedyropulik", "gaderypoluki")
	})
}) */
