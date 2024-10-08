package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMobileKeyboard(t *testing.T) {
	{
		result := MobileKeyboard("")
		var expect int = 0
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := MobileKeyboard("*#")
		var expect int = 2
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := MobileKeyboard("123")
		var expect int = 3
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := MobileKeyboard("codewars")
		var expect int = 26
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := MobileKeyboard("zruf")
		var expect int = 16
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := MobileKeyboard("thisisasms")
		var expect int = 37
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := MobileKeyboard("longwordwhichdontreallymakessense")
		var expect int = 107
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := MobileKeyboard("1234567890*#abcdefghijklmnopqrstuvwxyz")
		var expect int = 94
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
}

/*
	It("Basic Tests", func() {
		Expect(MobileKeyboard("")).To(Equal(0))
		Expect(MobileKeyboard("*#")).To(Equal(2))
		Expect(MobileKeyboard("123")).To(Equal(3))
		Expect(MobileKeyboard("codewars")).To(Equal(26))
		Expect(MobileKeyboard("zruf")).To(Equal(16))
		Expect(MobileKeyboard("thisisasms")).To(Equal(37))
		Expect(MobileKeyboard("longwordwhichdontreallymakessense")).To(Equal(107))
		Expect(MobileKeyboard("1234567890*#abcdefghijklmnopqrstuvwxyz")).To(Equal(94))
	})
*/
