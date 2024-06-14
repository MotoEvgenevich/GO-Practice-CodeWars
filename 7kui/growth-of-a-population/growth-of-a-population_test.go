package kata

import (
	"fmt"
	"testing"
)

type SolveTest struct {
	po       int
	percent  float64
	aug      int
	p        int
	Expected int
}

var tests = []SolveTest{
	{
		po:       1500000,
		percent:  2.5,
		aug:      10000,
		p:        2000000,
		Expected: 10,
	},
	{
		po:       1500000,
		percent:  0.25,
		aug:      1000,
		p:        2000000,
		Expected: 94,
	},
	{
		po:       1500000,
		percent:  0.25,
		aug:      1000,
		p:        2000000,
		Expected: 94,
	},
}

func TestNbYear(t *testing.T) {
	for _, test := range tests {
		output := NbYear(test.po, test.percent, test.aug, test.p)
		if output != test.Expected {
			t.Errorf("Expected output %v does not match expected value %v", output, test.Expected)
			continue
		}
	}
}

func BenchmarkNbYear(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NbYear(1500000, 0.25, 1000, 2000000)
	}
}

func ExampleNbYear() {
	fmt.Println(NbYear(1500000, 0.25, 1000, 2000000))
	// Output: 94
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
