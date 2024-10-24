package uniq

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniq(t *testing.T) {
	{
		result := Uniq([]string{"a", "a", "b", "b", "c", "a", "b", "c", "c"})
		expect := []string{"a", "b", "c", "a", "b", "c"}
		assert.Equal(t, result, expect, "не соответвует result & expect")
	}
	{
		result := Uniq([]string{"a", "a", "a", "b", "b", "b", "c", "c", "c"})
		expect := []string{"a", "b", "c"}
		assert.Equal(t, result, expect, "не соответвует result & expect")
	}
	{
		result := Uniq([]string{})
		expect := []string{}
		assert.Equal(t, result, expect, "не соответвует result & expect")
	}
	{
		result := Uniq([]string{"foo"})
		expect := []string{"foo"}
		assert.Equal(t, result, expect, "не соответвует result & expect")
	}
	{
		result := Uniq([]string{"bar"})
		expect := []string{"bar"}
		assert.Equal(t, result, expect, "не соответвует result & expect")
	}
	{
		result := Uniq([]string{""})
		expect := []string{""}
		assert.Equal(t, result, expect, "не соответвует result & expect")
	}
	{
		result := Uniq([]string{"a", "a"})
		expect := []string{"a"}
		assert.Equal(t, result, expect, "не соответвует result & expect")
	}
}

/*
	Expect(Uniq([]string{"a", "a", "b", "b", "c", "a", "b", "c", "c"})).To(Equal([]string{"a", "b", "c", "a", "b", "c"}))
	Expect(Uniq([]string{"a", "a", "a", "b", "b", "b", "c", "c", "c"})).To(Equal([]string{"a", "b", "c"}))
	Expect(Uniq([]string{})).To(Equal([]string{}))
	Expect(Uniq([]string{"foo"})).To(Equal([]string{"foo"}))
	Expect(Uniq([]string{"bar"})).To(Equal([]string{"bar"}))
	Expect(Uniq([]string{""})).To(Equal([]string{""}))
	Expect(Uniq([]string{"a", "a"})).To(Equal([]string{"a"}))
*/
