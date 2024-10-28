package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntRac(t *testing.T) {
	{
		result := IntRac(25, 1)
		expect := 4
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := IntRac(125348, 300)
		expect := 3
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := IntRac(125348981764, 356243)
		expect := 3
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := IntRac(236, 12)
		expect := 2
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := IntRac(48981764, 8000)
		expect := 3
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := IntRac(6999, 700)
		expect := 6
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := IntRac(16000, 400)
		expect := 5
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := IntRac(16000, 100)
		expect := 3
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := IntRac(2500, 60)
		expect := 2
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := IntRac(250000, 600)
		expect := 3
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}

}

/*
   dotest(25, 1, 4)
   dotest(125348, 300, 3)
   dotest(125348981764, 356243, 3)
   dotest(236, 12, 2)
   dotest(48981764, 8000, 3)
   dotest(6999, 700, 6)
   dotest(16000, 400, 5)
   dotest(16000, 100, 3)
   dotest(2500, 60, 2)
   dotest(250000, 600, 3)
   dotest(95367431640625, 1, 28)
   dotest(835871232077058, 1, 30)
   dotest(246391990316004, 1, 29)
   dotest(403832254158749, 1, 29)
   dotest(217414278071071, 1, 29)
   dotest(639593178334873, 1, 29)
   dotest(646895994940989, 1, 29)
   dotest(346147532294375, 1, 29)
   dotest(871397668454373, 1, 30)
   dotest(711652869429165, 1, 29)
   dotest(828061341209011, 1, 30)
*/
