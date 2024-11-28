package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainAllRots(t *testing.T) {
	{
		result := ContainAllRots("bsjq", []string{"bsjq", "qbsj", "sjqb", "twZNsslC", "jqbs"})
		expect := true
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := ContainAllRots("XjYABhR", []string{"TzYxlgfnhf", "yqVAuoLjMLy", "BhRXjYA", "YABhRXj", "hRXjYAB", "jYABhRX", "XjYABhR", "ABhRXjY"})
		expect := false
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := ContainAllRots("", []string{})
		expect := true
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
}

/*
   dotest("bsjq", []string{"bsjq", "qbsj", "sjqb", "twZNsslC", "jqbs"}, true)
   dotest("XjYABhR", []string{"TzYxlgfnhf", "yqVAuoLjMLy", "BhRXjYA", "YABhRXj", "hRXjYAB", "jYABhRX", "XjYABhR", "ABhRXjY"}, false)
   dotest("", []string{}, true)
*/
