package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToWeirdCase(t *testing.T) {
	{
		result := ToWeirdCase("abc def")
		expect := "AbC DeF"
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := ToWeirdCase("ABC")
		expect := "AbC"
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := ToWeirdCase("This is a test Looks like you passed")
		expect := "ThIs Is A TeSt LoOkS LiKe YoU PaSsEd"
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
}

/*
Expect(toWeirdCase("abc def")).To(Equal("AbC DeF"))
		Expect(toWeirdCase("ABC")).To(Equal("AbC"))
		Expect(toWeirdCase("This is a test Looks like you passed")).To(Equal("ThIs Is A TeSt LoOkS LiKe YoU PaSsEd"))
*/
