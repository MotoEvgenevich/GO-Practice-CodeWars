package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindLongestCommonPrefix(t *testing.T) {
	{
		actual := FindLongestCommonPrefix([]string{"flower", "flow", "flight"})
		expect := "fl"
		assert.Equal(t, expect, actual, "result doesn't equal expect")
	}
	{
		actual := FindLongestCommonPrefix([]string{"dog", "racecar", "car"})
		expect := ""
		assert.Equal(t, expect, actual, "result doesn't equal expect")
	}
}

/*
	Context("Find common prefix among multiple strings with same prefix", func() {
		It("Three strings with common prefix", func() {
			Expect(FindLongestCommonPrefix([]string{"abc123", "abcdef", "abc789"})).To(Equal("abc"))
		})

		It("Multiple strings with single character common prefix", func() {
			Expect(FindLongestCommonPrefix([]string{"a123", "a456", "a789"})).To(Equal("a"))
		})

		It("Multiple strings where common prefix is the entire shortest string", func() {
			Expect(FindLongestCommonPrefix([]string{"dog", "doghouse", "dogma"})).To(Equal("dog"))
		})

		It("Common prefix among multiple strings with same prefix", func() {
			Expect(FindLongestCommonPrefix([]string{"flower", "flow", "flight"})).To(Equal("fl"))
		})
	})

	Context("Find common prefix between two strings", func() {
		It("Two identical strings", func() {
			Expect(FindLongestCommonPrefix([]string{"hello", "hello"})).To(Equal("hello"))
		})

		It("Two strings with partial common prefix", func() {
			Expect(FindLongestCommonPrefix([]string{"world", "work"})).To(Equal("wor"))
		})

		It("Two strings with no common prefix", func() {
			Expect(FindLongestCommonPrefix([]string{"abc", "def"})).To(Equal(""))
		})
	})

	Context("Return empty string when no common prefix exists", func() {
		It("No common prefix", func() {
			Expect(FindLongestCommonPrefix([]string{"dog", "cat", "fish"})).To(Equal(""))
		})

		It("Strings starting with different characters", func() {
			Expect(FindLongestCommonPrefix([]string{"abc", "def", "ghi"})).To(Equal(""))
		})

		It("Two strings with no common prefix", func() {
			Expect(FindLongestCommonPrefix([]string{"abc", "def"})).To(Equal(""))
		})
	})
})

var _ = Describe("Edge Cases", func() {
	Context("Handles single string in slice", func() {
		It("should return the single string", func() {
			Expect(FindLongestCommonPrefix([]string{"hello"})).To(Equal("hello"))
		})
	})

	Context("Handles strings of different lengths", func() {
		It("should return the longest common prefix for different lengths", func() {
			Expect(FindLongestCommonPrefix([]string{"flower", "flow", "flight"})).To(Equal("fl"))
		})
	})

	Context("Handles strings with special characters", func() {
		It("should handle special characters", func() {
			Expect(FindLongestCommonPrefix([]string{"$%^&*", "$%^&*abc"})).To(Equal("$%^&*"))
		})
	})

	Context("Handles strings with mixed case letters", func() {
		It("should handle mixed case letters", func() {
			Expect(FindLongestCommonPrefix([]string{"Hello", "HeLlO", "hEllO"})).To(Equal("hello"))
		})
	})

	Context("Handles strings with unicode characters", func() {
		It("should handle unicode characters", func() {
			Expect(FindLongestCommonPrefix([]string{"café", "caféaulait"})).To(Equal("café"))
		})
	})

	Context("Handles strings containing only numbers", func() {
		It("should handle numeric strings", func() {
			Expect(FindLongestCommonPrefix([]string{"123", "12345", "123456"})).To(Equal("123"))
		})
	})
})
*/
