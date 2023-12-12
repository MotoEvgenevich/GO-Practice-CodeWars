package kata

import (
	"fmt"
	"testing"
)

type SolveTest struct {
	Str      string
	Expected string
}

var tests = []SolveTest{
	{
		Str:      "alpha beta beta gamma gamma gamma delta alpha beta beta gamma gamma gamma delta",
		Expected: "alpha beta gamma delta",
	},
	{
		Str:      "my cat is my cat fat",
		Expected: "my cat is fat",
	},
}

func TestRemoveDuplicateWords(t *testing.T) {
	for _, test := range tests {
		output := RemoveDuplicateWords(test.Str)
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

func BenchmarkRemoveDuplicateWords(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RemoveDuplicateWords("alpha beta beta gamma gamma gamma delta alpha beta beta gamma gamma gamma delta")
	}
}

func ExampleRemoveDuplicateWords() {
	fmt.Println(RemoveDuplicateWords("alpha beta beta gamma gamma gamma delta alpha beta beta gamma gamma gamma delta"))
	// Output: alpha beta gamma delta
}

//To start coverage test use this commands in terminal:
//go test -coverprofile=coverage.out
//go tool cover -html=coverage.out

//To Start Benchmark Test use this commands:
//go test -bench=.
//go test -bench=<NAMEFUNCTION>

/*
TEST CACES:
/*
	"alpha beta beta gamma gamma gamma delta alpha beta beta gamma gamma gamma delta" "alpha beta gamma delta"
	"my cat is my cat fat"   "my cat is fat"
*/
