package kata

import (
	"fmt"
	"testing"
)

type SolveTest struct {
	Word     string
	Expected string
}

var tests = []SolveTest{
	{
		Word:     "hello",
		Expected: "0.1.2.2.3",
	},
	{
		Word:     "HeLlo",
		Expected: "0.1.2.2.3",
	},
	{
		Word:     "helLo",
		Expected: "0.1.2.2.3",
	},
	{
		Word:     "Hippopotomonstrosesquippedaliophobia",
		Expected: "0.1.2.2.3.2.3.4.3.5.3.6.7.4.8.3.7.9.7.10.11.1.2.2.9.12.13.14.1.3.2.0.3.15.1.13",
	},
}

func TestWordPattern(t *testing.T) {
	for _, test := range tests {
		output := WordPattern(test.Word)
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

func BenchmarkWordPattern(b *testing.B) {
	for i := 0; i < b.N; i++ {
		WordPattern("Hippopotomonstrosesquippedaliophobia")
	}
}

func ExampleWordPattern() {
	fmt.Println(WordPattern("hello"))
	// Output: 0.1.2.2.3
}

//To start coverage test use this commands in terminal:
//go test -coverprofile=coverage.out
//go tool cover -html=coverage.out

//To Start Benchmark Test use this commands:
//go test -bench=.
//go test -bench=<NAMEFUNCTION>

/*
TEST CACES:
        dotest("hello", "0.1.2.2.3")
        dotest("heLlo", "0.1.2.2.3")
        dotest("helLo", "0.1.2.2.3")
        dotest("Hippopotomonstrosesquippedaliophobia", "0.1.2.2.3.2.3.4.3.5.3.6.7.4.8.3.7.9.7.10.11.1.2.2.9.12.13.14.1.3.2.0.3.15.1.13")
*/
