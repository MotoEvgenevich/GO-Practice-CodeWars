package kata

import (
	"fmt"
	"testing"
)

type SolveTest struct {
	Limit    int
	Expected int
}

var tests = []SolveTest{
	{
		Limit:    8,
		Expected: 10,
	},
	{
		Limit:    111111,
		Expected: 60696,
	},
	{
		Limit:    8675309,
		Expected: 4613732,
	},
}

func TestSumEvenFibonacci(t *testing.T) {
	for _, test := range tests {
		output := SumEvenFibonacci(test.Limit)
		if output != test.Expected {
			t.Errorf("Expected of output %v does not match expected value %v", output, test.Expected)
		}
	}
}

func BenchmarkSumEvenFibonacci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumEvenFibonacci(8675309)
	}
}

func ExampleSumEvenFibonacci() {
	fmt.Println(SumEvenFibonacci(8675309))
	// Output: 4613732
}

//To start coverage test use this commands in terminal:
//go test -coverprofile=coverage.out
//go tool cover -html=coverage.out

//To Start Benchmark Test use this commands:
//go test -bench=.
//go test -bench=<NAMEFUNCTION>

/* var _ = Describe("SumEvenFibonacci", func() {
	It("should return 10 for input 8", func() {
	  Expect(SumEvenFibonacci(8)).To(Equal(10))
	})
	It("should return 60696 for input 111111", func() {
	  Expect(SumEvenFibonacci(111111)).To(Equal(60696))
	})
	It("should return 4613732 for input 8675309", func() {
	  Expect(SumEvenFibonacci(8675309)).To(Equal(4613732))
	})
  }) */
