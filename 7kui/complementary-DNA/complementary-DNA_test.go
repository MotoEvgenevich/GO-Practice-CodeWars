package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDNAStrand(t *testing.T) {
	{
		result := DNAStrand("AAAA")
		expect := "TTTT"
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := DNAStrand("ATTGC")
		expect := "TAACG"
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := DNAStrand("GTAT")
		expect := "CATA"
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
}

/*
   Expect(DNAStrand("AAAA")).To(Equal("TTTT"))
   Expect(DNAStrand("ATTGC")).To(Equal("TAACG"))
   Expect(DNAStrand("GTAT")).To(Equal("CATA"))
*/
