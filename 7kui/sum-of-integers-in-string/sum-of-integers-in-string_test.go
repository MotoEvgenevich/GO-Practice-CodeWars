package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumOfIntegersInString(t *testing.T) {

	{
		result := SumOfIntegersInString("The30quick20brown10f0x1203jumps914ov3r1349the102l4zy dog")
		expect := 3635
		assert.Equal(t, result, expect, "value of result doesn't equal expect")

	}
	{
		result := SumOfIntegersInString("14.2")
		expect := 16
		assert.Equal(t, result, expect, "value of result doesn't equal expect")

	}

}
