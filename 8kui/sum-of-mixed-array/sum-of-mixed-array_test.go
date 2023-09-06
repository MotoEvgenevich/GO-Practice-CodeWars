package kata

import (
	"fmt"
	"testing"
)

type SolveData struct {
	Arr      []any
	Expected int
}

var testData = []SolveData{
	{
		Arr:      []any{"5", "0", 9, 3, 2, 1, "9", 6, 7},
		Expected: 42,
	},
	{
		Arr:      []any{9, 3, "7", "3"},
		Expected: 22,
	},
	{
		Arr:      []any{"3", 6, 6, 0, "5", 8, 5, "6", 2, "0"},
		Expected: 41,
	},
	{
		Arr:      []any{"1", "5", "8", 8, 9, 9, 2, "3"},
		Expected: 45,
	},
	{
		Arr:      []any{8, 0, 0, 8, 5, 7, 2, 3, 7, 8, 6, 7},
		Expected: 61,
	},
}

func TestSumMix(t *testing.T) {
	for _, test := range testData {
		output := SumMix(test.Arr)
		if output != test.Expected {
			t.Errorf("value of output %v does not match expected value %v", output, test.Expected)
			continue
		}
	}
}

func BenchmarkSumMix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumMix([]any{8, 0, 0, 8, 5, 7, 2, 3, 7, 8, 6, 7})
	}
}

func ExampleSumMix() {
	fmt.Println(SumMix([]any{8, 0, 0, 8, 5, 7, 2, 3, 7, 8, 6, 7}))
	// Output: 61
}

//To start coverage test use this commands in terminal:
//go test -coverprofile=coverage.out
//go tool cover -html=coverage.out

//To Start Benchmark Test use this commands:
//go test -bench=.
//go test -bench=<NAMEFUNCTION>

/* doTest([]any{9, 3, "7", "3"}, 22)
doTest([]any{"5", "0", 9, 3, 2, 1, "9", 6, 7}, 42)
doTest([]any{"3", 6, 6, 0, "5", 8, 5, "6", 2,"0"}, 41)
doTest([]any{"1", "5", "8", 8, 9, 9, 2, "3"}, 45)
doTest([]any{8, 0, 0, 8, 5, 7, 2, 3, 7, 8, 6, 7}, 61) */
