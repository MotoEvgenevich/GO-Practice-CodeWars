package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBandNameGenerator(t *testing.T) {
	{
		result := bandNameGenerator("knife")
		expect := "The Knife"
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := bandNameGenerator("tart")
		expect := "Tartart"
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := bandNameGenerator("sandles")
		expect := "Sandlesandles"
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := bandNameGenerator("bed")
		expect := "The Bed"
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
}

/*
		It("should return the correct values", func() {
		Expect(bandNameGenerator(string("knife"))).Should(BeEquivalentTo("The Knife"))
		Expect(bandNameGenerator(string("tart"))).Should(BeEquivalentTo("Tartart"))
		Expect(bandNameGenerator(string("sandles"))).Should(BeEquivalentTo("Sandlesandles"))
		Expect(bandNameGenerator(string("bed"))).Should(BeEquivalentTo("The Bed"))
	})
})
*/
