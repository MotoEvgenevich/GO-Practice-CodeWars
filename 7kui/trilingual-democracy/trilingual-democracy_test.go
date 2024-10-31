package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrilingualDemocracy(t *testing.T) {
	{
		result := TrilingualDemocracy("FFF")
		expect := "F"
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := TrilingualDemocracy("IIK")
		expect := "K"
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := TrilingualDemocracy("DFK")
		expect := "I"
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}

}

/*
		Expect(TrilingualDemocracy("FFF")).To(Equal("F"))
	})
	It("IIK", func() {
		Expect(TrilingualDemocracy("IIK")).To(Equal("K"))
	})
	It("DFK", func() {
		Expect(TrilingualDemocracy("DFK")).To(Equal("I"))
*/
