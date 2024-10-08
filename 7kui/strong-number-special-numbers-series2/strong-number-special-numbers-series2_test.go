package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrongFunc(t *testing.T) {
	t.Run("Test Strong", func(t *testing.T) {
		assert.Equal(t, "STRONG!!!!", Strong(1))
		assert.Equal(t, "STRONG!!!!", Strong(2))
		assert.Equal(t, "STRONG!!!!", Strong(145))
		assert.Equal(t, "Not Strong !!", Strong(7))
		assert.Equal(t, "Not Strong !!", Strong(93))
		assert.Equal(t, "Not Strong !!", Strong(185))
	})
}

/*
   Tester(1, "STRONG!!!!")
   Tester(2, "STRONG!!!!")
   Tester(145, "STRONG!!!!")
   Tester(7, "Not Strong !!")
   Tester(93, "Not Strong !!")
   Tester(185, "Not Strong !!")
*/
