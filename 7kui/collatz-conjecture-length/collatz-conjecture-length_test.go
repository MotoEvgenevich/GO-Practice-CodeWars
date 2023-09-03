package kata

import (
	"fmt"
	"testing"
)

type SolveTest struct {
	Num      int
	Expected int
}

var tests = []SolveTest{
	{
		Num:      20,
		Expected: 8,
	},
	{
		Num:      15,
		Expected: 18,
	},
	{
		Num:      100,
		Expected: 26,
	},
	{
		Num:      10,
		Expected: 7,
	},
	{
		Num:      500,
		Expected: 111,
	},
	{
		Num:      1,
		Expected: 1,
	},
	{
		Num:      1000000000,
		Expected: 101,
	},
	{
		Num:      1000000000000000,
		Expected: 276,
	},
}

func TestCollatz(t *testing.T) {
	for _, test := range tests {
		output := Collatz(test.Num)
		if output != test.Expected {
			t.Errorf("Expected output %v does not match expected value %v", output, test.Expected)
			continue
		}
	}
}

func BenchmarkCollatz(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Collatz(1000000000000000)
	}
}

func ExampleCollatz() {
	fmt.Println(Collatz(20))
	// Output: 8
}

//To start coverage test use this commands in terminal:
//go test -coverprofile=coverage.out
//go tool cover -html=coverage.out

//To Start Benchmark Test use this commands:
//go test -bench=.
//go test -bench=<NAMEFUNCTION>

/* dotest(20, 8)
dotest(15, 18)
dotest(100, 26)
dotest(10, 7)
dotest(500, 111)
dotest(1, 1)
dotest(1000000000, 101)
dotest(1000000000000000, 276) */
