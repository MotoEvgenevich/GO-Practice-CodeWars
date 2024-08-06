package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxRot(t *testing.T) {
	{
		expected := int64(85821534)
		result := MaxRot(38458215)
		assert.Equal(t, expected, result, "не соотвествуют данным:  dotest(38458215, 85821534)")
	}
	{
		expected := int64(988103115)
		result := MaxRot(195881031)
		assert.Equal(t, expected, result, "не соотвествуют данным:  dotest(195881031, 988103115)")
	}
	{
		expected := int64(962193428)
		result := MaxRot(896219342)
		assert.Equal(t, expected, result, "не соотвествуют данным:  dotest(195881031, 988103115)")
	}

}

/*
   dotest(38458215, 85821534)
   dotest(195881031, 988103115)
   dotest(896219342, 962193428)
*/
