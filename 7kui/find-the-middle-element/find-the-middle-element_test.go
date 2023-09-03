package kata

import (
	"fmt"
	"testing"
)

type SolveTest struct {
	Array    [3]int
	Expected int
}

var tests = []SolveTest{
	{
		Array:    [3]int{2, 3, 1},
		Expected: 0,
	},
	{
		Array:    [3]int{5, 10, 14},
		Expected: 1,
	},
	{
		Array:    [3]int{1, 3, 4},
		Expected: 1,
	},
	{
		Array:    [3]int{15, 10, 14},
		Expected: 2,
	},
	{
		Array:    [3]int{-4, -23, 4},
		Expected: 0,
	},
	{
		Array:    [3]int{-15, -10, 14},
		Expected: 1,
	},
}

func TestGimme(t *testing.T) {
	for _, test := range tests {
		output := Gimme(test.Array)
		if output != test.Expected {
			t.Errorf("Expected output %v does not match expected value %v", output, test.Expected)
			continue
		}
	}
}

func BenchmarkGimme(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Gimme([3]int{-15, -10, 14})
	}
}

func ExampleGimme() {
	fmt.Println(Gimme([3]int{-15, -10, 14}))
	// Output: 1
}

//To start coverage test use this commands in terminal:
//go test -coverprofile=coverage.out
//go tool cover -html=coverage.out

//To Start Benchmark Test use this commands:
//go test -bench=.
//go test -bench=<NAMEFUNCTION>

/* dotest([3]int{2, 3, 1}, 0)
dotest([3]int{5, 10, 14}, 1)
dotest([3]int{1, 3, 4}, 1)
dotest([3]int{15, 10, 14}, 2)
dotest([3]int{-4, -23, 4}, 0)
dotest([3]int{-15, -10, 14}, 1) */
