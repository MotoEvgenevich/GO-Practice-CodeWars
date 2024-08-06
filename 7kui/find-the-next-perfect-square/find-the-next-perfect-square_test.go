package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindNextSquare(t *testing.T) {
	{
		result := FindNextSquare(121)
		expected := int64(144)
		assert.Equal(t, result, expected, "данные не равны в тесте: 121")
	}
	{
		result := FindNextSquare(625)
		expected := int64(676)
		assert.Equal(t, result, expected, "данные не равны в тесте: 625")
	}
	{
		result := FindNextSquare(319225)
		expected := int64(320356)
		assert.Equal(t, result, expected, "данные не равны в тесте: 319225")
	}
	{
		result := FindNextSquare(15241383936)
		expected := int64(15241630849)
		assert.Equal(t, result, expected, "данные не равны в тесте: 15241383936")
	}
	{
		result := FindNextSquare(155)
		expected := int64(-1)
		assert.Equal(t, result, expected, "данные не равны в тесте: 155")
	}
}

/*
   Expect(FindNextSquare(int64(121))).To(Equal(int64(144)))
   Expect(FindNextSquare(int64(625))).To(Equal(int64(676)))
   Expect(FindNextSquare(int64(319225))).To(Equal(int64(320356)))
   Expect(FindNextSquare(int64(15241383936))).To(Equal(int64(15241630849)))
   Expect(FindNextSquare(int64(155))).To(Equal(int64(-1)))
*/
