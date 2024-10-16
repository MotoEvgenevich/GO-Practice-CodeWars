package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAreCoprime(t *testing.T) {
	{
		result := AreCoprime(20, 27)
		expect := true
		assert.Equal(t, result, expect, "result doesn't equal expect")

	}

	{
		result := AreCoprime(12, 39)
		expect := false
		assert.Equal(t, result, expect, "result doesn't equal expect")

	}
}

/*
	     It("Sample tests", func() {
       Expect(AreCoprime(20, 27)).To(BeTrue())
       Expect(AreCoprime(12, 39)).To(BeFalse())
     })
*/
