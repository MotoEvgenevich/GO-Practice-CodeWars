package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCats(t *testing.T) {
	{
		result := Cats(1, 5)
		var expect int = 2
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := Cats(1, 1)
		var expect int = 0
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := Cats(2, 5)
		var expect int = 1
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := Cats(2, 4)
		var expect int = 2
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
}

/*
  It( "Fixed tests", func() {
    Expect(Cats(1,5)).To(Equal(2));
    Expect(Cats(1,1)).To(Equal(0));
    Expect(Cats(2,5)).To(Equal(1));
    Expect(Cats(2,4)).To(Equal(2), "Mew");
  });
*/
