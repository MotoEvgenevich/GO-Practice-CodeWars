package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFixStringCase(t *testing.T) {
	{
		result := solve("a")
		expect := "a"
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := solve("Z")
		expect := "Z"
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := solve("coDe")
		expect := "code"
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := solve("CODe")
		expect := "CODE"
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := solve("coDE")
		expect := "code"
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := solve("coDE")
		expect := "code"
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := solve("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
		expect := "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz"
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
}
func TestBestPractice(t *testing.T) {
	{
		result := solve("a")
		expect := "a"
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := solve("Z")
		expect := "Z"
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := solve("coDe")
		expect := "code"
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := solve("CODe")
		expect := "CODE"
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := solve("coDE")
		expect := "code"
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := solve("coDE")
		expect := "code"
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := solve("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
		expect := "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz"
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
}

/*
		Expect(solve("a")).To(Equal("a"))
		Expect(solve("Z")).To(Equal("Z"))
		Expect(solve("coDe")).To(Equal("code"))
		Expect(solve("CODe")).To(Equal("CODE"))
		Expect(solve("coDE")).To(Equal("code"))
		Expect(solve("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")).To(Equal("abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz"))
	})
*/
