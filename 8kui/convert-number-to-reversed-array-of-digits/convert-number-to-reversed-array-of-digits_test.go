package kata

import (
	"fmt"
	"testing"
)

type SolveData struct {
	Num      int
	Expected []int
}

var testData = []SolveData{
	{
		Num:      9876543210,
		Expected: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
	},
	{
		Num:      0,
		Expected: []int{0},
	},
	{
		Num:      35231,
		Expected: []int{1, 3, 2, 5, 3},
	},
}

func TestDigitize(t *testing.T) {
	for _, test := range testData {
		output := Digitize(test.Num)
		if len(output) != len(test.Expected) {
			t.Errorf("Expected length of output %v does not match expected length %v", len(output), len(test.Expected))

		}
		for i, v := range output {
			if v != test.Expected[i] {
				t.Errorf("Output %v does not match Expected %v", output, test.Expected)
			}
		}
	}
}

func BenchmarkDigitize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Digitize(9876543210)
	}
}

func ExampleDigitize() {
	fmt.Println(Digitize(9876543210))
	// Output: [0 1 2 3 4 5 6 7 8 9]
}

//To start coverage test use this commands in terminal:
//go test -coverprofile=coverage.out
//go tool cover -html=coverage.out

//To Start Benchmark Test use this commands:
//go test -bench=.
//go test -bench=<NAMEFUNCTION>
