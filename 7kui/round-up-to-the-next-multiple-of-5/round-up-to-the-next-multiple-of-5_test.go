package kata

import (
	"fmt"
	"testing"
)

type SolveTest struct {
	N        int
	Expected int
}

var tests = []SolveTest{
	{
		N:        0,
		Expected: 0,
	},
	{
		N:        1,
		Expected: 5,
	},
	{
		N:        -1,
		Expected: 0,
	},
	{
		N:        5,
		Expected: 5,
	},
	{
		N:        7,
		Expected: 10,
	},
	{
		N:        20,
		Expected: 20,
	},
	{
		N:        39,
		Expected: 40,
	},
	{
		N:        121,
		Expected: 125,
	},
	{
		N:        555,
		Expected: 555,
	},
}

func TestRoundToNext5(t *testing.T) {
	for _, test := range tests {
		output := RoundToNext5(test.N)
		if output != test.Expected {
			t.Errorf("Expected of output %v does not match expected value %v", output, test.Expected)
		}
	}
}

func BenchmarkRoundToNext5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RoundToNext5(20)
	}
}

func ExampleRoundToNext5() {
	fmt.Println(RoundToNext5(554))
	// Output: 555
}

//To start coverage test use this commands in terminal:
//go test -coverprofile=coverage.out
//go tool cover -html=coverage.out

//To Start Benchmark Test use this commands:
//go test -bench=.
//go test -bench=<NAMEFUNCTION>

/* [][2]int{{0, 0}, {1, 5}, {-1, 0}, {5, 5}, {7, 10}, {20, 20}, {39, 40}, {121, 125}, {555,555}}  */
