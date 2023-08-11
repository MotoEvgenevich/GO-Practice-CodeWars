package kata

import (
	"fmt"
	"testing"
)

type SolveData struct {
	Num      int
	Expected string
}

var testData = []SolveData{
	{
		Num:      4,
		Expected: "1 sheep...2 sheep...3 sheep...4 sheep...",
	},
	{
		Num:      6,
		Expected: "1 sheep...2 sheep...3 sheep...4 sheep...5 sheep...6 sheep...",
	},
	{
		Num:      10,
		Expected: "1 sheep...2 sheep...3 sheep...4 sheep...5 sheep...6 sheep...7 sheep...8 sheep...9 sheep...10 sheep...",
	},
	{
		Num:      1,
		Expected: "1 sheep...",
	},
	{
		Num:      3,
		Expected: "1 sheep...2 sheep...3 sheep...",
	},
}

func TestCountSheep(t *testing.T) {
	for _, test := range testData {
		output := CountSheep(test.Num)
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

func BenchmarkCountSheep(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CountSheep(4)
	}
}

func ExampleCountSheep() {
	fmt.Println(CountSheep(4))
	// Output: 1 sheep...2 sheep...3 sheep...4 sheep...
}

//To start coverage test use this commands in terminal:
//go test -coverprofile=coverage.out
//go tool cover -html=coverage.out

//To Start Benchmark Test use this commands:
//go test -bench=.
//go test -bench=<NAMEFUNCTION>
