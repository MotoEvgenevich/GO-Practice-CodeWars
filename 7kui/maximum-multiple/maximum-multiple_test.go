package kata

import (
	"fmt"
	"testing"
)

type SolveTest struct {
	Divisible int
	Bound     int
	Expected  int
}

var tests = []SolveTest{
	{
		Divisible: 2,
		Bound:     7,
		Expected:  6,
	},
	{
		Divisible: 3,
		Bound:     10,
		Expected:  9,
	},
	{
		Divisible: 7,
		Bound:     17,
		Expected:  14,
	},
	{
		Divisible: 10,
		Bound:     50,
		Expected:  50,
	},
	{
		Divisible: 37,
		Bound:     200,
		Expected:  185,
	},
	{
		Divisible: 7,
		Bound:     100,
		Expected:  98,
	},
}

func TestSolve(t *testing.T) {
	for _, test := range tests {
		output := MaxMultiple(test.Divisible, test.Bound)
		if output != test.Expected {
			t.Errorf("Expected int %v does not match of output %v", output, test.Expected)

		}
	}
}

func BenchmarkMaxMultiple(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MaxMultiple(37, 200)
	}
}

func ExampleMaxMultiple() {
	fmt.Println(MaxMultiple(2, 7))
	// Output: 6
}

//To start coverage test use this commands in terminal:
//go test -coverprofile=coverage.out
//go tool cover -html=coverage.out

//To Start Benchmark Test use this commands:
//go test -bench=.
//go test -bench=<NAMEFUNCTION>

/* It("MaxMultiple(2, 7)", func() { Expect(MaxMultiple(2, 7)).To(Equal(6)) })
   It("MaxMultiple(3, 10)", func() { Expect(MaxMultiple(3, 10)).To(Equal(9)) })
   It("MaxMultiple(7, 17)", func() { Expect(MaxMultiple(7, 17)).To(Equal(14)) })
   It("MaxMultiple(10, 50)", func() { Expect(MaxMultiple(10, 50)).To(Equal(50)) })
   It("MaxMultiple(37, 200)", func() { Expect(MaxMultiple(37, 200)).To(Equal(185)) })
   It("MaxMultiple(7, 100)", func() { Expect(MaxMultiple(7, 100)).To(Equal(98)) }) */
