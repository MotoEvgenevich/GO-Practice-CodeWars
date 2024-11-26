package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAvg(t *testing.T) {
	{
		result := NewAvg([]float64{14.0, 30.0, 5.0, 7.0, 9.0, 11.0, 16.0}, 90)
		expect := int64(628)
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := NewAvg([]float64{14, 30, 5, 7, 9, 11, 15}, 92)
		expect := int64(645)
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := NewAvg([]float64{1400.25, 30000.76, 5.56, 7, 9, 11, 15.48, 120.98}, 2000)
		expect := int64(-1)
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}

}

/*
   var a = []float64{14.0, 30.0, 5.0, 7.0, 9.0, 11.0, 16.0}
   dotest(a, 90, 628)
   a = []float64{14, 30, 5, 7, 9, 11, 15}
   dotest(a, 92, 645)
   a = []float64{1400.25, 30000.76, 5.56, 7, 9, 11, 15.48, 120.98}
   dotest(a, 2000, -1)
*/
