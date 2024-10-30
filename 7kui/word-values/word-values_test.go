package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNameValue(t *testing.T) {
	{
		result := NameValue([]string{"abc", "abc", "abc", "abc"})
		expect := []int{6, 12, 18, 24}
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}

	{
		result := NameValue([]string{"abc", "abc abc"})
		expect := []int{6, 24}
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
}

/*
Input:
t1 := []string{"abc","abc","abc","abc"}
s1 := []int{6,12,18,24}
t2 := []string{"codewars","abc","xyz"}
s2 := []int{88,12,225}

*/
