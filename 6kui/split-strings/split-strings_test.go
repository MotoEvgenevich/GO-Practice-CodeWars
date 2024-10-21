package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	{
		result := Solution("abc")
		expect := []string{"ab", "c_"}
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := Solution("dawsd")
		expect := []string{"da", "ws", "d_"}
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := Solution("awsaws")
		expect := []string{"aw", "sa", "ws"}
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}

}

/*
   Expect(Solution("abc")).To(Equal([]string{"ab", "c_"}))
   Expect(Solution("dawsd")).To(Equal([]string{"da", "ws", "d_"}))
   Expect(Solution("awsaws")).To(Equal([]string{"aw", "sa", "ws"}))
*/
