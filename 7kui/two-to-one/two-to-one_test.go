package kata

import (
	"fmt"
	"testing"
)

type SolveTest struct {
	S1       string
	S2       string
	Expected string
}

var tests = []SolveTest{
	{
		S1:       "aretheyhere",
		S2:       "yestheyarehere",
		Expected: "aehrsty",
	},
	{
		S1:       "loopingisfunbutdangerous",
		S2:       "lessdangerousthancoding",
		Expected: "abcdefghilnoprstu",
	},
}

func TestTwoToOne(t *testing.T) {
	for _, test := range tests {
		output := TwoToOne(test.S1, test.S2)
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

func BenchmarkTwoToOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TwoToOne("aretheyhere", "yestheyarehere")
	}
}

func ExampleTwoToOne() {
	fmt.Println(TwoToOne("aretheyhere", "yestheyarehere"))
	// Output: aehrsty
}

//To start coverage test use this commands in terminal:
//go test -coverprofile=coverage.out
//go tool cover -html=coverage.out

//To Start Benchmark Test use this commands:
//go test -bench=.
//go test -bench=<NAMEFUNCTION>

/*
Expect(TwoToOne("aretheyhere", "yestheyarehere")).To(Equal("aehrsty"))
Expect(TwoToOne("loopingisfunbutdangerous", "lessdangerousthancoding")).To(Equal("abcdefghilnoprstu"))
Examples:
*/
