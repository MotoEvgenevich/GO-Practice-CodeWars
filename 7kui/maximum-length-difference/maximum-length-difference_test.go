package kata

import (
	"fmt"
	"testing"
)

type SolveTest struct {
	A1       []string
	A2       []string
	Expected int
}

var tests = []SolveTest{
	{
		A1:       []string{"ejjjjmmtthh", "zxxuueeg", "aanlljrrrxx", "dqqqaaabbb", "oocccffuucccjjjkkkjyyyeehh"},
		A2:       []string{"bbbaaayddqbbrrrv"},
		Expected: 10,
	},
	{
		A1:       []string{"dsfsdfds", "sadasd"},
		A2:       []string{},
		Expected: -1,
	},
}

func TestMxDifLg(t *testing.T) {
	for _, test := range tests {
		output := MxDifLg(test.A1, test.A2)
		if output != test.Expected {
			t.Errorf("Expected value of output %v does not match expected value %v", output, test.Expected)
			continue
		}

	}
}

func BenchmarkMxDifLg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MxDifLg([]string{"ejjjjmmtthh", "zxxuueeg", "aanlljrrrxx", "dqqqaaabbb", "oocccffuucccjjjkkkjyyyeehh"}, []string{"bbbaaayddqbbrrrv"})
	}
}

func ExampleMxDifLg() {
	fmt.Println(MxDifLg([]string{"ejjjjmmtthh", "zxxuueeg", "aanlljrrrxx", "dqqqaaabbb", "oocccffuucccjjjkkkjyyyeehh"}, []string{"bbbaaayddqbbrrrv"}))
	// Output: 10
}

//To start coverage test use this commands in terminal:
//go test -coverprofile=coverage.out
//go tool cover -html=coverage.out

//To Start Benchmark Test use this commands:
//go test -bench=.
//go test -bench=<NAMEFUNCTION>

/*
s1 = []string{"ejjjjmmtthh", "zxxuueeg", "aanlljrrrxx", "dqqqaaabbb", "oocccffuucccjjjkkkjyyyeehh"}
s2 = []string{"bbbaaayddqbbrrrv"}
*/
