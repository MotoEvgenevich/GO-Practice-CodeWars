package kata

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type SolveData struct {
	Str      string
	Expected string
}

var testData = []SolveData{
	{
		Str:      "most trees are blue",
		Expected: "Most Trees Are Blue",
	},
	{
		Str:      "All the rules in this world were made by someone no smarter than you. So make your own.",
		Expected: "All The Rules In This World Were Made By Someone No Smarter Than You. So Make Your Own.",
	},
	{
		Str:      "When I die. then you will realize",
		Expected: "When I Die. Then You Will Realize",
	},
	{
		Str:      "Jonah Hill is a genius",
		Expected: "Jonah Hill Is A Genius",
	},
	{
		Str:      "Dying is mainstream",
		Expected: "Dying Is Mainstream",
	},
}

func TestToJadenCaseLength(t *testing.T) {
	{
		result := ToJadenCase("All the rules in this world were made by someone no smarter than you. So make your own.")
		expect := "All The Rules In This World Were Made By Someone No Smarter Than You. So Make Your Own."
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := ToJadenCase("When I die. then you will realize")
		expect := "When I Die. Then You Will Realize"
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := ToJadenCase("Jonah Hill is a genius")
		expect := "Jonah Hill Is A Genius"
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}

}
func TestToJadenCase(t *testing.T) {
	for _, test := range testData {
		output := ToJadenCase(test.Str)
		if len(output) != len(test.Expected) {
			t.Errorf("Expected length of output %v does not match expected length %v", len(output), len(test.Expected))
			continue
		}
		for i := 0; i < len(output); i++ {
			if output[i] != test.Expected[i] {
				t.Errorf("Output %v at index %d does not match expected %v", output[i], i, test.Expected[i])
			}
		}
	}
}

func BenchmarkToJadenCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ToJadenCase("All the rules in this world were made by someone no smarter than you. So make your own.")
	}
}

func ExampleToJadenCase() {
	fmt.Println(ToJadenCase("All the rules in this world were made by someone no smarter than you. So make your own."))
	// Output: All The Rules In This World Were Made By Someone No Smarter Than You. So Make Your Own.
}

//To start coverage test use this commands in terminal:
//go test -coverprofile=coverage.out
//go tool cover -html=coverage.out

//To Start Benchmark Test use this commands:
//go test -bench=.
//go test -bench=<NAMEFUNCTION>

/* Expect(ToJadenCase("most trees are blue")).To(Equal("Most Trees Are Blue"))
Expect(ToJadenCase("All the rules in this world were made by someone no smarter than you. So make your own.")).To(Equal("All The Rules In This World Were Made By Someone No Smarter Than You. So Make Your Own."))
Expect(ToJadenCase("When I die. then you will realize")).To(Equal("When I Die. Then You Will Realize"))
Expect(ToJadenCase("Jonah Hill is a genius")).To(Equal("Jonah Hill Is A Genius"))
Expect(ToJadenCase("Dying is mainstream")).To(Equal("Dying Is Mainstream")) */
